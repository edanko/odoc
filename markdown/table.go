package markdown

import (
	"strings"
)

type Table struct {
	cols int
	body [][]string
}

func NewTable(header []string) *Table {
	t := new(Table)

	t.cols = len(header)
	t.body = append(t.body, header)

	var sep []string
	for i := 0; i < t.cols; i++ {
		sep = append(sep, "----")
	}
	t.body = append(t.body, sep)

	return t
}

func (t *Table) AppendRow(content []string) {
	if len(content) != t.cols {
		for i := 0; i < t.cols-len(content); i++ {
			content = append(content, " ")
		}
	}
	t.body = append(t.body, content)
}

func (t *Table) String() string {
	var buffer strings.Builder
	for _, row := range t.body {
		buffer.WriteString("|")
		for _, col := range row {
			buffer.WriteString(col)
			buffer.WriteString("|")
		}
		buffer.WriteString("\n")

	}
	return buffer.String()
}
