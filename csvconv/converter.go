package csvconv

import (
	"io"
)

// CSVConverter is an interface that can be used as any CSV converter
type CSVConverter interface {
	ConvertFromCSV(csvFile io.Reader) string
	ConvertToCSV(targetFile io.Reader) string
}
