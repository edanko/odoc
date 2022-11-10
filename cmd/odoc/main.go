package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/edanko/odoc/format"
	"github.com/edanko/odoc/oapi"
)

func main() {

	// filePath := "./openapi.yaml"
	filePath := "./openapi2.yaml"
	// filePath := "./petstore-expanded.yaml"

	specs, err := oapi.Parse(context.Background(), filePath)
	if err != nil {
		panic(err)
	}

	formatter := format.New()

	formatter.Header(2, "Paths")

	for _, spec := range specs {
		formatter.Paragraph(format.Bold(spec.Method) + " " + spec.Path)
		if spec.Description != "" {
			formatter.Paragraph(spec.Description)
		}

		formatter.Header(3, "Request parameters")
		if spec.Request.ContentType != "" {
			formatter.Paragraph("Content type: " + spec.Request.ContentType)
		}

		if !spec.Request.PathParams.IsEmpty() {
			formatter.Paragraph("Path params")
			formatter.Paragraph(PropertiesToTable(spec.Request.PathParams))
		}
		if !spec.Request.QueryParams.IsEmpty() {
			formatter.Paragraph("query")
			formatter.Paragraph(PropertiesToTable(spec.Request.QueryParams))
		}
		if !spec.Request.Headers.IsEmpty() {
			formatter.Paragraph("headers")
			formatter.Paragraph(PropertiesToTable(spec.Request.Headers))
		}
		if !spec.Request.Body.IsEmpty() {
			formatter.Paragraph("Body")
			formatter.Paragraph(PropertiesToTable(spec.Request.Body))
		}

		formatter.Header(3, "Responses")
		head := strconv.Itoa(spec.Response.StatusCode) + " " + http.StatusText(spec.Response.StatusCode)
		if spec.Response.Description != "" {
			head += ": " + spec.Response.Description
		}
		formatter.Paragraph(head)

		formatter.Paragraph("Content type: " + spec.Response.ContentType)

		if !spec.Response.Headers.IsEmpty() {
			formatter.Paragraph("headers")
			formatter.Paragraph(PropertiesToTable(spec.Response.Headers))
		}
		if !spec.Response.Body.IsEmpty() {
			formatter.Paragraph("body")
			formatter.Paragraph(PropertiesToTable(spec.Response.Body))
		}

		formatter.Paragraph("---")

	}

	f, err := os.Create("DOCS.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(formatter.String())
	if err != nil {
		panic(err)
	}
}

func PropertiesToTable(properties []oapi.Property) string {
	if len(properties) == 0 {
		return ""
	}

	table := format.NewTable([]string{
		"Required",
		"Name",
		"In",
		"Type / Format",
		"Description",
		"Enum",
	})

	for _, param := range properties {
		if param.Type != "object" {

			row := get(param)
			table.AppendRow(row)

		}
		for _, child := range param.Children {
			row := get(child)
			table.AppendRow(row)
		}
	}
	return table.String()
}

func get(param oapi.Property) []string {
	var req string
	if param.Required {
		req = "*"
	} else {
		req = " "
	}
	row := []string{
		req,
		param.Name,
		param.In,
		param.Type + "/" + param.SubType + "/" + param.Format,
		param.Description,
		strings.Join(param.Enum, ", "),
	}
	return row
}
