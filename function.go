package gopherjson

import "fmt"

// CustomFunction represents a function that can be serialized and deserialized.
type CustomFunction struct {
	FunctionName string
	// Additional fields specific to your use case
}

func (cf CustomFunction) Serialize() interface{} {
	return cf.FunctionName
}

func (cf *CustomFunction) Deserialize(value interface{}) error {
	str, ok := value.(string)

	if !ok {
		return fmt.Errorf("value is not a string")
	}

	cf.FunctionName = str // Adjust deserialization logic according to your use case

	return nil
}

func (cf *CustomFunction) UnmarshalJSON(data []byte) error {
	str := string(data)
	cf.FunctionName = str // Adjust deserialization logic according to your use case
	return nil
}
