package goldtest

import (
	"encoding/json"
	"flag"
	"fmt"

	"testing"

	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files")

// GetGoldenFile - return byte array. Allows you to use golden file content as you want.
// Also, if you pass flag "-update", golden file will be updated (or created).
func GetGoldenFile(actual []byte, fileName string) ([]byte, error) {
	if *update {
		if err := writeFile(fileName, actual); err != nil {
			return nil, err
		}
	}

	byteBody, err := readFile(fileName)

	if err != nil {
		return nil, err
	}

	return byteBody, nil
}

// Assert - allow testing any type of data. Function is able to load,
// update or create golden file and compare it with provided interface.
// Bytes for goldenfiles are created by converting interface to string, then to byte array.
func Assert(t *testing.T, actual interface{}, fileName string) {
	actualBytes := []byte(fmt.Sprintf("%v", actual.(interface{})))

	goldenBytes, err := GetGoldenFile(actualBytes, fileName)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, actualBytes, goldenBytes)
}

// AssertJSON - allow testing any type of data. Function is able to load,
// update or create golden file and compare it with provided interface.
// Bytes for goldenfiles are created by marshaling provided daya.
func AssertJSON(t *testing.T, actual interface{}, fileName string) {
	actualBytes, err := marshal(actual)

	if err != nil {
		t.Fatal(err)
	}

	goldenBytes, err := GetGoldenFile(actualBytes, fileName)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, actualBytes, goldenBytes)
}

func marshal(dataToMarshal interface{}) ([]byte, error) {
	bytes, err := json.MarshalIndent(dataToMarshal, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
