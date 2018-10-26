package alexa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"
)

func Test_unmarshalling(t *testing.T) {
	file := loadTestFile("alexa_request.json", func() {
		t.Fatal("Unable to read testdata file")
	})

	var request Request
	if e := json.Unmarshal(file, &request); e != nil {
		t.Error("Unable to parse json")
	}

	resolutions := request.Body.Intent.Slots["SLOT_NAME"].Resolutions.ResolutionPerAuthority
	values := resolutions[0].Values

	if len(values) != 2 {
		t.Errorf("Number of values expected 2, but was %d", len(values))
	}

	checkValue := func(index int) {
		expectedValue := fmt.Sprintf(`SLOT_VALUE_%s`, strconv.Itoa(index))

		if values[index].Value.Name != expectedValue {
			t.Errorf("Expected `%s` but was `%s` \n", expectedValue, values[index].Value.Name)
		}

		expectedID := strconv.Itoa(index)
		if values[index].Value.ID != expectedID {
			t.Errorf("Expected `%s` but was `%s` \n", expectedID, values[index].Value.ID)
		}
	}

	checkValue(0)
	checkValue(1)
}

func loadTestFile(name string, failureHandler func()) []byte {
	path := filepath.Join("testdata", name)
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		failureHandler()
	}

	return bytes
}
