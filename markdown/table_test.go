package markdown

import (
	"reflect"
	"testing"
)

func TestNewTable(t *testing.T) {
	type args struct {
		header []string
	}
	tests := []struct {
		name string
		args args
		want *Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTable(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_AppendRow(t1 *testing.T) {
	type fields struct {
		cols int
		body [][]string
	}
	type args struct {
		content []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				cols: tt.fields.cols,
				body: tt.fields.body,
			}
			t.AppendRow(tt.args.content)
		})
	}
}

func TestTable_String(t1 *testing.T) {
	type fields struct {
		cols int
		body [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				cols: tt.fields.cols,
				body: tt.fields.body,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
