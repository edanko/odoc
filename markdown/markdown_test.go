package markdown

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBold(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "successfully return bold text",
			text: "test",
			want: "**test**",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bold(tt.text)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMarkdown_CodeBlock(t *testing.T) {
	type fields struct {
		builder strings.Builder
	}
	type args struct {
		lang string
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Markdown{
				builder: tt.fields.builder,
			}
			m.CodeBlock(tt.args.lang, tt.args.code)
		})
	}
}

func TestMarkdown_Header(t *testing.T) {
	tests := []struct {
		name  string
		level int
		text  string
		want  string
	}{
		{
			name:  "return level 1 header from text",
			level: 1,
			text:  "test",
			want:  "# test\n\n",
		},
		{
			name:  "return level 2 header from text",
			level: 2,
			text:  "test",
			want:  "## test\n\n",
		},
		{
			name:  "return level 3 header from text",
			level: 3,
			text:  "test",
			want:  "### test\n\n",
		},
		{
			name:  "return level 4 header from text",
			level: 4,
			text:  "test",
			want:  "#### test\n\n",
		},
		{
			name:  "return level 5 header from text",
			level: 5,
			text:  "test",
			want:  "##### test\n\n",
		},
		{
			name:  "return level 6 header from text",
			level: 6,
			text:  "test",
			want:  "###### test\n\n",
		},
		{
			name:  "return level 6 header from text with other levels",
			level: 16,
			text:  "test",
			want:  "###### test\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New()
			m.Header(tt.level, tt.text)

			require.Equal(t, tt.want, m.String())
		})
	}
}

func TestMarkdown_Link(t *testing.T) {
	type fields struct {
		builder strings.Builder
	}
	type args struct {
		text string
		href string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Markdown{
				builder: tt.fields.builder,
			}
			m.Link(tt.args.text, tt.args.href)
		})
	}
}

func TestMarkdown_Paragraph(t *testing.T) {
	type fields struct {
		builder strings.Builder
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Markdown{
				builder: tt.fields.builder,
			}
			m.Paragraph(tt.args.text)
		})
	}
}

func TestMarkdown_String(t *testing.T) {
	type fields struct {
		builder strings.Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Markdown{
				builder: tt.fields.builder,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarkdown_Text(t *testing.T) {
	type fields struct {
		builder strings.Builder
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Markdown{
				builder: tt.fields.builder,
			}
			m.Text(tt.args.text)
		})
	}
}
