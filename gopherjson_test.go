package gopherjson

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func TestSerializeDeserialize(t *testing.T) {
	// Define a test data structure
	type TestData struct {
		ID          int
		Name        string
		Date        CustomDate
		Regex       CustomRegex
		CustomFunc  CustomFunction
		SubTestData []TestData
	}

	// Create a test instance of the data structure
	now := time.Now()
	data := TestData{
		ID:   1,
		Name: "John",
		Date: CustomDate{now},
		Regex: CustomRegex{
			regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
		},
		CustomFunc: CustomFunction{
			FunctionName: "myFunction",
		},
		SubTestData: []TestData{
			{
				ID:   2,
				Name: "Jane",
				Date: CustomDate{now.AddDate(0, 0, 1)},
				Regex: CustomRegex{
					regexp.MustCompile(`^[A-Z]+$`),
				},
				CustomFunc: CustomFunction{
					FunctionName: "anotherFunction",
				},
			},
		},
	}

	// Serialize the test data
	serialized, err := Serialize(data)
	if err != nil {
		t.Errorf("Failed to serialize: %v", err)
	}

	// Deserialize the serialized data
	deserialized, err := Deserialize(serialized)
	if err != nil {
		t.Errorf("Failed to deserialize: %v", err)
	}

	// Compare the original and deserialized data
	if !reflect.DeepEqual(data, deserialized) {
		t.Errorf("Deserialized data does not match the original")
	}

	// Print the results
	fmt.Println("Original Data:", data)
	fmt.Println("Serialized Data:", serialized)
	fmt.Println("Deserialized Data:", deserialized)
}
