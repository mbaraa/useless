package useless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode"
)

// StringsExtended has string operations that are not in the "strings" module
type StringsExtended struct{}

// NewExtendedStrings returns a new StringsExtended instance
func NewExtendedStrings() *StringsExtended {
	return &StringsExtended{}
}

// GetStringArrayFromJSONBytes is same as GetStringArrayFromJSON
// but it receives a byte array that has the JSON data instead
func (s *StringsExtended) GetStringArrayFromJSONBytes(jsonBytes []byte, arrayName ...string) ([]string, error) {
	jsonReader := bytes.NewReader(jsonBytes)
	return s.GetStringArrayFromJSON(jsonReader, arrayName...)
}

// GetStringArrayFromJSON returns a string array with a specific name from a json file
// if an array name is provided it returns that specific array, otherwise the whole json is an array :)
func (s *StringsExtended) GetStringArrayFromJSON(jsonFile io.Reader, arrayName ...string) ([]string, error) {
	var finalArray []string // hmm
	var err error

	if arrayName != nil { // the array is an object from the JSON file
		finalArray, err = s.getNamedJSONArray(jsonFile, arrayName[0])
		if err != nil {
			return nil, err
		}

	} else { // the whole JSON file is a JSON array
		finalArray, err = s.getJSONArray(jsonFile)
		if err != nil {
			return nil, err
		}
	}

	// happily ever after
	return finalArray, nil
}

// getNamedJSONArray returns a string array from the given json reader with the given array name
func (s *StringsExtended) getNamedJSONArray(jsonFile io.Reader, arrayName string) (finalArray []string, err error) {
	jMap := make(map[string]interface{})

	err = json.NewDecoder(jsonFile).Decode(&jMap)
	if err != nil {
		return nil, err
	}

	// convert []interface{} to []string
	finalArray = strings.Split(
		strings.TrimRight(
			strings.TrimLeft(fmt.Sprint(jMap[arrayName]), "["),
			"]"),
		" ",
	)

	// happily ever after
	return
}

// getJSONArray returns a string array that is the whole json file
func (s StringsExtended) getJSONArray(jsonFile io.Reader) (finalArray []string, err error) {
	jData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jData, &finalArray)
	if err != nil {
		return nil, err
	}

	// happily ever after
	return
}

// MakeSpongeBobCase returns the given in SpOnGeBoBcAsE
// also it receives an optional flag to start string with lower case, default value = false
func (s *StringsExtended) MakeSpongeBobCase(str string, startLower ...bool) string {
	spongeBob, sl := "", false

	if startLower != nil {
		sl = startLower[0]
	}

	var case1, case2 int
	if sl { // lower first
		case1, case2 = 1, 0
	} else { // lower first
		case1, case2 = 0, 1
	}

	for i := range str {
		if i%2 == 0 {
			spongeBob += string(unicode.To(case1, rune(str[i])))
		} else {
			spongeBob += string(unicode.To(case2, rune(str[i])))
		}
	}

	return spongeBob
}
