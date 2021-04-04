package csvconv

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// MarkdownConverter is a CSV, Md bidirectional converter
type MarkdownConverter struct{}

// NewMarkdownConverter returns a csv <-> md converter
func NewMarkdownConverter() *MarkdownConverter {
	return &MarkdownConverter{}
}

// ConvertFromCSV converts a given csv file into an md formatted string
// with the data from the csv file by adding elements between pipes
// md is amazing isn't it??
func (m *MarkdownConverter) ConvertFromCSV(csvFile io.Reader) string {

	columnsTitles, csvFields, err := getCSVData(csvFile)
	if err != nil {
		return ""
	}

	md := m.getTablePrefix(columnsTitles)
	numColumns := len(columnsTitles)
	// adding entries
	for i, entry := range csvFields {

		if (i+1+(numColumns-1))%numColumns == 0 { // beginning entry
			md += fmt.Sprintf("\n| %s |", strings.Trim(entry, " "))

		} else if (i+1)%numColumns == 0 { // ending entry
			md += fmt.Sprintf("%s |", strings.Trim(entry, " "))

		} else { // middle entry
			md += fmt.Sprintf(" %s | ", strings.Trim(entry, " "))
		}
	}

	return md
}

// ConvertToCSV converts a given md file into a csv string
// by splitting the md string into lines then splitting the lines by pipes
// then appending each entry into the resulting csv string
func (m *MarkdownConverter) ConvertToCSV(markdownFile io.Reader) string {
	md, readErr := ioutil.ReadAll(markdownFile)
	if readErr != nil {
		return ""
	}
	// final csv string
	csv := ""

	lines := strings.Split(string(md), "\n")

	var sep uint8 = 44 // comma is the default separator

	for j, line := range lines {
		rowEntries := strings.Split(
			strings.Trim(line, "|"), // strip pipes from each line to avoid extra commas :)
			"|")
		for i, entry := range rowEntries {
			if i >= len(rowEntries)-1 { // when the last row element is reached flip the comma into a space
				sep = 32 // a space is really harmless,
				// also since it's getting rid of additional commas it's a superior
			}

			if entry != "" && !strings.Contains(entry, "---") { // skip blank strings & titles separators
				csv += fmt.Sprintf("%s%c ", strings.Trim(entry, " "), sep) // strip from spaces
			}
		}
		if j > 0 { // skip line separator line
			csv += "\n"
		}
		sep = 44
	}

	return csv
}

// getTablePrefix returns md table prefix with the given titles
func (m *MarkdownConverter) getTablePrefix(columnsTitles []string) string {
	prefix := "|"

	// columns titles
	for _, title := range columnsTitles {
		prefix += fmt.Sprintf(" %s |", title)
	}
	prefix += "\n"

	// adding separators between titles and the other entries
	prefix += "| :--- |"                                       // first column
	prefix += strings.Repeat(" :---: |", len(columnsTitles)-2) // middle columns
	prefix += "---: |"                                         // last column

	return prefix
}
