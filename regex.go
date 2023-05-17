package gopherjson

import (
	"fmt"
	"regexp"
)

// CustomRegex represents a regular expression that can be serialized and deserialized.
type CustomRegex struct {
	*regexp.Regexp
}

func (cr CustomRegex) Serialize() interface{} {
	return cr.String()
}

func (cr *CustomRegex) Deserialize(value interface{}) error {
	str, ok := value.(string)

	if !ok {
		return fmt.Errorf("value is not a string")
	}

	regex, err := regexp.Compile(str)

	if err != nil {
		return fmt.Errorf("error deserializing CustomRegex: %v", err)
	}

	cr.Regexp = regex

	return nil
}

func (cr *CustomRegex) UnmarshalJSON(data []byte) error {
	str := string(data)
	regex, err := regexp.Compile(str)

	if err != nil {
		return fmt.Errorf("error unmarshaling CustomRegex: %v", err)
	}

	cr.Regexp = regex

	return nil
}
