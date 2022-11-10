package oapi

import (
	"context"
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name         string
		filepath     string
		wantHeader   *Header
		wantSpecsLen int
		wantErr      bool
	}{
		{
			name:     "successfully parse valid file",
			filepath: "testdata/petstore-expanded.yaml",
			wantHeader: &Header{
				Title:       "Swagger Petstore",
				Description: "A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification",
			},
			wantSpecsLen: 7,
			wantErr:      false,
		},
		{
			name:     "return error on non-existing file",
			filepath: "not-exist.yaml",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotSpecs, err := Parse(context.Background(), tt.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, tt.wantHeader, gotHeader)
			require.Len(t, gotSpecs, tt.wantSpecsLen)
		})
	}
}

func TestProperties_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		ps   Properties
		want bool
	}{
		{
			name: "return true on empty properties",
			ps:   Properties{Property{}},
			want: true,
		},
		{
			name: "return false on non-empty properties",
			ps: Properties{
				Property{
					Required:    true,
					Name:        "some property name",
					Description: "some description",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractHeaders(t *testing.T) {
	type args struct {
		headers openapi3.Headers
	}
	tests := []struct {
		name string
		args args
		want Properties
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseAPISpec(t *testing.T) {
	type args struct {
		method string
		path   string
		op     *openapi3.Operation
	}
	tests := []struct {
		name string
		args args
		want []*APISpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseAPISpec(tt.args.method, tt.args.path, tt.args.op); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAPISpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseResponseStatus(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "return int(200) on input string(\"200\")", args: args{status: "200"}, want: 200},
		{name: "return int(500) on input string(\"500\")", args: args{status: "500"}, want: 500},
		{name: "return int(200) on invalid input value", args: args{status: "200.5"}, want: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseResponseStatus(tt.args.status); got != tt.want {
				t.Errorf("parseResponseStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schemaToProperty(t *testing.T) {
	type args struct {
		name   string
		in     string
		schema *openapi3.Schema
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := schemaToProperty(tt.args.name, tt.args.in, tt.args.schema); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("schemaToProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
