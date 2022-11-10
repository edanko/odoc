package oapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/samber/lo"
)

type Header struct {
	Title       string
	Description string
}

type APISpec struct {
	Description string
	Path        string
	Method      string
	Request     Request
	Response    Response
}

type Request struct {
	ContentType string
	PathParams  Properties
	QueryParams Properties
	Headers     Properties
	Body        Properties
}

type Response struct {
	ContentType string
	StatusCode  int
	Description string
	Headers     Properties
	Body        Properties
}

type Property struct {
	Required    bool
	Name        string
	Description string
	Type        string
	SubType     string
	Enum        []string
	In          string
	Format      string
	Children    []Property
}

func (p Property) IsEmpty() bool {
	return reflect.DeepEqual(p, Property{})
}

type Properties []Property

func (ps Properties) IsEmpty() bool {
	for _, p := range ps {
		if !p.IsEmpty() {
			return false
		}
	}
	return true
}

func Parse(ctx context.Context, filePath string) (*Header, []*APISpec, error) {
	loader := &openapi3.Loader{
		Context:               ctx,
		IsExternalRefsAllowed: true,
	}

	var doc *openapi3.T
	u, err := url.Parse(filePath)
	if err == nil && u.Scheme != "" && u.Host != "" {
		doc, err = loader.LoadFromURI(u)
		if err != nil {
			return nil, nil, err
		}
	} else {
		doc, err = loader.LoadFromFile(filePath)
		if err != nil {
			return nil, nil, err
		}
	}

	header := &Header{
		Title:       doc.Info.Title,
		Description: doc.Info.Description,
	}

	var specs []*APISpec
	for path, pathItem := range doc.Paths {
		specs = append(specs, parseAPISpec(http.MethodDelete, path, pathItem.Delete)...)
		specs = append(specs, parseAPISpec(http.MethodGet, path, pathItem.Get)...)
		specs = append(specs, parseAPISpec(http.MethodPost, path, pathItem.Post)...)
		specs = append(specs, parseAPISpec(http.MethodPut, path, pathItem.Put)...)
		specs = append(specs, parseAPISpec(http.MethodPatch, path, pathItem.Patch)...)
	}

	return header, specs, nil
}

func parseAPISpec(method string, path string, op *openapi3.Operation) []*APISpec {
	specs := make([]*APISpec, 0)
	if op == nil {
		return nil
	}
	reqContent := make(map[string]*openapi3.MediaType)
	if op.RequestBody != nil && op.RequestBody.Value != nil {
		reqContent = op.RequestBody.Value.Content
	}
	reqHeaders := make(Properties, 0)
	queryParams := make(Properties, 0)
	pathParams := make(Properties, 0)
	for _, param := range op.Parameters {
		if param.Value.Name == "" {
			param.Value.Name = param.Ref
		}
		if param.Value.Schema != nil {
			property := schemaToProperty(param.Value.Name, param.Value.In, param.Value.Schema.Value)
			if param.Value.In == "path" {
				pathParams = append(pathParams, property)
			} else if param.Value.In == "header" {
				reqHeaders = append(reqHeaders, property)
			} else {
				queryParams = append(queryParams, property)
			}
		}
	}

	for status, resp := range op.Responses {
		respHeaders := extractHeaders(resp.Value.Headers)

		for resContentType, resMedia := range resp.Value.Content {
			if len(reqContent) > 0 {
				for reqContentType, reqMedia := range reqContent {
					spec := &APISpec{
						Description: op.Description,
						Method:      method,
						Path:        path,
						Request: Request{
							Headers:     reqHeaders,
							QueryParams: queryParams,
							PathParams:  pathParams,
							Body:        make(Properties, 0),
							ContentType: reqContentType,
						},
						Response: Response{
							Headers:     respHeaders,
							ContentType: resContentType,
							Description: *resp.Value.Description,
							Body:        Properties{schemaToProperty("", "body", resMedia.Schema.Value)},
							StatusCode:  parseResponseStatus(status),
						},
					}
					if op.RequestBody != nil && op.RequestBody.Value != nil {
						spec.Request.ContentType = reqContentType
						spec.Request.Body = append(spec.Request.Body,
							schemaToProperty("", "body", reqMedia.Schema.Value))
					}
					specs = append(specs, spec)
				}
			} else {
				spec := &APISpec{
					Description: op.Description,
					Method:      method,
					Path:        path,
					Request: Request{
						Headers:     reqHeaders,
						QueryParams: queryParams,
						PathParams:  pathParams,
						Body:        make(Properties, 0),
						ContentType: "",
					},
					Response: Response{
						Headers:     respHeaders,
						ContentType: resContentType,
						Body:        Properties{schemaToProperty("", "body", resMedia.Schema.Value)},
						StatusCode:  parseResponseStatus(status),
					},
				}
				specs = append(specs, spec)
			}
		}
	}
	return specs
}

func schemaToProperty(name string, in string, schema *openapi3.Schema) Property {
	property := Property{
		Required:    lo.Contains(schema.Required, name),
		Name:        name,
		Description: schema.Description,
		Type:        schema.Type,
		Format:      schema.Format,
		In:          in,
		Children:    make(Properties, 0),
	}
	if schema.Items != nil {
		property.SubType = schema.Items.Value.Type
	}
	if schema.Enum != nil {
		property.Enum = make([]string, len(schema.Enum))
		for i, next := range schema.Enum {
			property.Enum[i] = fmt.Sprintf("%v", next)
		}
	}
	if schema.Items != nil && schema.Items.Value != nil {
		for name, prop := range schema.Items.Value.Properties {
			property.Children = append(property.Children, schemaToProperty(name, in, prop.Value))
		}
	}
	for name, prop := range schema.Properties {
		property.Children = append(property.Children, schemaToProperty(name, in, prop.Value))
	}
	if schema.AdditionalProperties != nil {
		for name, prop := range schema.AdditionalProperties.Value.Properties {
			property.Children = append(property.Children, schemaToProperty(name, in, prop.Value))
		}
	}
	return property
}

func extractHeaders(headers openapi3.Headers) Properties {
	var res Properties
	for k, header := range headers {
		if header.Value.Schema == nil {
			continue
		}
		if header.Value.Name == "" {
			header.Value.Name = k
		}
		property := schemaToProperty(header.Value.Name, header.Value.In, header.Value.Schema.Value)
		res = append(res, property)
	}
	return res
}

func parseResponseStatus(status string) int {
	code, err := strconv.Atoi(status)
	if err != nil {
		code = 200
	}
	return code
}
