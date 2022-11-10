package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/edanko/odoc/markdown"
	"github.com/edanko/odoc/oapi"
)

func main() {
	var includeHeader bool
	flag.BoolVar(&includeHeader, "header", false, "include header in the output")

	flag.Parse()

	inputPath := flag.Arg(0)
	if inputPath == "" {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage:\n\t%s [option...] <input.yaml or url> [output]\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	header, specs, err := oapi.Parse(context.Background(), inputPath)
	if err != nil {
		const format = "cannot parse file: %w"
		log.Fatalln(fmt.Errorf(format, err))
	}

	md := markdown.New()

	if includeHeader {
		md.Header(1, header.Title)
		md.Paragraph(header.Description)
	}

	md.Header(2, "Paths")

	for _, spec := range specs {
		md.Paragraph(markdown.Bold(spec.Method) + " " + spec.Path)
		if spec.Description != "" {
			md.Paragraph(spec.Description)
		}

		md.Header(3, "Request parameters")
		if spec.Request.ContentType != "" {
			md.Paragraph("Content type: " + spec.Request.ContentType)
		}

		if !spec.Request.PathParams.IsEmpty() {
			md.Paragraph("Path params")
			md.Paragraph(PropertiesToTable(spec.Request.PathParams))
		}
		if !spec.Request.QueryParams.IsEmpty() {
			md.Paragraph("query")
			md.Paragraph(PropertiesToTable(spec.Request.QueryParams))
		}
		if !spec.Request.Headers.IsEmpty() {
			md.Paragraph("headers")
			md.Paragraph(PropertiesToTable(spec.Request.Headers))
		}
		if !spec.Request.Body.IsEmpty() {
			md.Paragraph("Body")
			md.Paragraph(PropertiesToTable(spec.Request.Body))
		}

		md.Header(3, "Responses")
		head := strconv.Itoa(spec.Response.StatusCode) + " " + http.StatusText(spec.Response.StatusCode)
		if spec.Response.Description != "" {
			head += ": " + spec.Response.Description
		}
		md.Paragraph(head)

		md.Paragraph("Content type: " + spec.Response.ContentType)

		if !spec.Response.Headers.IsEmpty() {
			md.Paragraph("headers")
			md.Paragraph(PropertiesToTable(spec.Response.Headers))
		}
		if !spec.Response.Body.IsEmpty() {
			md.Paragraph("body")
			md.Paragraph(PropertiesToTable(spec.Response.Body))
		}

		md.Paragraph("---")

	}

	outputPath := flag.Arg(1)
	if outputPath == "" {
		fmt.Println(md.String())
		return
	}

	f, err := os.Create(outputPath)
	if err != nil {
		const format = "cannot open output file: %w"
		log.Fatalln(fmt.Errorf(format, err))
	}
	defer f.Close()

	_, err = f.WriteString(md.String())
	if err != nil {
		const format = "cannot write string to file: %w"
		log.Fatalln(fmt.Errorf(format, err))
	}
}

func PropertiesToTable(properties []oapi.Property) string {
	if len(properties) == 0 {
		return ""
	}

	table := markdown.NewTable([]string{
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
