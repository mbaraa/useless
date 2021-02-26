package useless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// StringsExtended has string operations that are not in the "strings" module
type StringsExtended struct{}

// NewExtendedStrings returns a new StringsExtended instance
func NewExtendedStrings() *StringsExtended {
	return &StringsExtended{}
}

// GetStringArrayFromJSONBytes same as GetStringArrayFromJSON but it receives a byte array with the JSON data instead
func (s *StringsExtended) GetStringArrayFromJSONBytes(jsonBytes []byte, arrayName ...string) ([]string, error) {
	jsonReader := bytes.NewReader(jsonBytes)
	return s.GetStringArrayFromJSON(jsonReader, arrayName...)
}

// GetStringArrayFromJSON returns a string array with a specific name from a json file
func (s *StringsExtended) GetStringArrayFromJSON(jsonFile io.Reader, arrayName ...string) ([]string, error) {
	var finalArray []string // hmm

	if arrayName != nil { // the array is an object from the JSON file
		jMap := make(map[string]interface{})

		err := json.NewDecoder(jsonFile).Decode(&jMap)
		if err != nil {
			return nil, err
		}

		// convert []interface{} to []string
		finalArray = strings.Split(
			strings.TrimRight(
				strings.TrimLeft(fmt.Sprint(jMap[arrayName[0]]), "["),
				"]"),
			" ",
		)

	} else { // the whole JSON file is a JSON array

		jData, err := ioutil.ReadAll(jsonFile)
		err = json.Unmarshal(jData, &finalArray)
		if err != nil {
			return nil, err
		}
	}

	// happily ever after
	return finalArray, nil
}
