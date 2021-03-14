package csvconv

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// HTMLConverter is a CSV, HTML bidirectional converter
type HTMLConverter struct{}

// NewHTMLConverter returns a csv <-> html converter
func NewHTMLConverter() *HTMLConverter {
	return &HTMLConverter{}
}

// ConvertFromCSV converts a given csv file into an html formatted document
// that has a table containing the data from the csv file, by magic well no
// by putting all the csv data inside a nice looking html table
func (h *HTMLConverter) ConvertFromCSV(csvFile io.Reader) string {
	columnsTitles, csvFields, err := getCSVData(csvFile)
	if err != nil {
		panic(err)
	}

	// html page
	html := h.getPagePrefix()
	html += h.getTablePrefix(columnsTitles)

	numColumns := len(columnsTitles)

	for i := range csvFields {
		html += fmt.Sprintf("\t<td>%s</td>\n", csvFields[i])
		if (i+1)%numColumns == 0 && (i+1) >= numColumns {
			html += "</tr><tr>\n"
		}
	}

	html += h.getTablePostfix()
	html += h.getPagePostfix()

	return html
}

// ConvertToCSV converts a given html file into a csv string
// by splitting the html string by "<t" ie the start of <tr>, <td>, or <th>
// also no other tags has 'r', 'd', or 'h' as the second char in it
// that after the closing triangular parenthesis exist some actual data
func (h *HTMLConverter) ConvertToCSV(htmlFile io.Reader) string {
	html, readErr := ioutil.ReadAll(htmlFile)
	if readErr != nil {
		panic(readErr)
	}
	// final csv string
	titles := h.getTableTitles(string(html)) // titles :)
	csv := strings.Join(titles, ",")

	fields := strings.Split(string(html), "<t")
	// track row's elements number, to avoid adding additional commas :)
	rowElementsCount := 1
	var sep uint8 = ',' // comma is the default separator

	for i, v := range fields {

		if rowElementsCount >= len(titles) { // when the last row element is reached flip the comma into a space
			sep = ' ' // a space is really harmless,
			// also since it's getting rid of additional commas it's really amazing
		}

		if v[0] == 'd' { // ie <td> ie a table column
			csv += fmt.Sprintf("%s%c",
				v[strings.Index(v, ">")+1:  // first occurrence of '>'
				strings.Index(v, "</td>")], // end of field's name
				sep)

			rowElementsCount++ // lol

		} else if i > 2 && v[0] == 'r' { // ie <th> ie a table row
			csv += "\n"
			rowElementsCount = 1 // reset row elements counter
			sep = ','            // reset separator
		}
	}

	return csv
}

// getTableTitles returns every* table title that can be found in the html document
// every title: well this is still an alpha, I'll fix that later I swear :)
func (*HTMLConverter) getTableTitles(html string) []string {
	var titles []string

	fields := strings.Split(html, "<t")
	for _, v := range fields {
		if v[0] == 'h' { // ie <th> ie column's title
			titles = append(titles, v[strings.Index(v, ">")+1: // first occurrence of '>'
			strings.Index(v, "</th>")])                        // end of field's name
		}
	}

	return titles
}

func (*HTMLConverter) getPagePrefix() string {
	return `
<!DOCTYPE html>
<html>
<head></head>
<body>
`
}

func (*HTMLConverter) getPagePostfix() string {
	return `
</body>
</html>`
}

func (h *HTMLConverter) getTablePrefix(columnsTitles []string) string {
	table := `
<table border=1 fram=hsides rules=rows,columns>
<tr>`

	for i := 0; i < len(columnsTitles); i++ {
		table += fmt.Sprintf("<th scope=\"col\">%s</th>\n", columnsTitles[i])
	}

	table += `
</tr>
<tr> `

	return table
}

func (*HTMLConverter) getTablePostfix() string {
	return `
</tr>
</table>`
}
