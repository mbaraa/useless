package useless

import "github.com/baraa-almasri/useless/csvconv"

// NewMarkdownConverter returns a csv <-> md converter
func NewMarkdownConverter() *csvconv.MarkdownConverter {
	return &csvconv.MarkdownConverter{}
}

// NewHTMLConverter returns a csv <-> html converter
func NewHTMLConverter() *csvconv.HTMLConverter {
	return &csvconv.HTMLConverter{}
}
