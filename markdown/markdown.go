package markdown

import (
	"fmt"
	"strings"
)

type Markdown struct {
	builder strings.Builder
}

func New() *Markdown {
	return &Markdown{
		builder: strings.Builder{},
	}
}

func (m *Markdown) String() string {
	return m.builder.String()
}

func (m *Markdown) Header(level int, text string) {
	switch level {
	case 1:
		m.builder.WriteString(fmt.Sprintf("# %s\n\n", text))
	case 2:
		m.builder.WriteString(fmt.Sprintf("## %s\n\n", text))
	case 3:
		m.builder.WriteString(fmt.Sprintf("### %s\n\n", text))
	case 4:
		m.builder.WriteString(fmt.Sprintf("#### %s\n\n", text))
	case 5:
		m.builder.WriteString(fmt.Sprintf("##### %s\n\n", text))
	default:
		m.builder.WriteString(fmt.Sprintf("###### %s\n\n", text))
	}
}

func (m *Markdown) Link(text, href string) {
	if text == "" {
		return
	}
	if href == "" {
		m.builder.WriteString(text)
	}

	m.builder.WriteString(fmt.Sprintf("[%s](%s)", text, href))
}

func (m *Markdown) Text(text string) {
	m.builder.WriteString(text)
}

func (m *Markdown) Paragraph(text string) {
	if text == "" {
		return
	}
	m.builder.WriteString(text + "\n\n")
}

func (m *Markdown) CodeBlock(lang, code string) {
	m.builder.WriteString(fmt.Sprintf("```%s\n%s\n```\n\n", lang, strings.TrimSpace(code)))
}

func Bold(text string) string {
	return "**" + text + "**"
}
