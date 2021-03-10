package csvconv

import (
	"encoding/csv"
	"io"
)

// getCSVData juices every little thing from a csv file :)
// ie it returns csv columns titles and fields, and returns
// an error indicates file read error, or wrong number of columns
func getCSVData(csvFile io.Reader) ([]string, []string, error) {
	// get columns titles
	csvReader := csv.NewReader(csvFile)
	columnsTitles, readErr := csvReader.Read()
	if readErr != nil {
		return nil, nil, readErr
	}

	// check for wrong number of columns
	csvData, columnsErr := csvReader.ReadAll()
	if columnsErr != nil {
		return nil, nil, columnsErr
	}

	// happily ever after :)
	return columnsTitles, getCSVDataList(csvData[0:]), nil
}

// getCSVDataList returns a list of a given csv 2dim string slice
func getCSVDataList(csv [][]string) []string {
	var finalCSV []string

	for _, s := range csv {
		for _, s2 := range s {
			finalCSV = append(finalCSV, s2)
		}
	}

	return finalCSV
}
