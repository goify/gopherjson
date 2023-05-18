package gopherjson

import (
	"reflect"
	"testing"
	"time"
)

// Define a sample struct for testing
type Person struct {
	Name     string     `json:"name"`
	Age      int        `json:"age"`
	Birthday CustomDate `json:"birthday"`
}

func TestSerializeAndDeserialize(t *testing.T) {
	// Create a sample object
	p := Person{
		Name:     "John Doe",
		Age:      30,
		Birthday: CustomDate{time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	// Serialize the object
	serialized, err := Serialize(p)
	if err != nil {
		t.Fatalf("Error serializing object: %v", err)
	}

	// Deserialize the serialized value
	var deserialized Person
	err = Deserialize(serialized, &deserialized)
	if err != nil {
		t.Fatalf("Error deserializing value: %v", err)
	}

	// Compare the original and deserialized objects
	if !reflect.DeepEqual(p, deserialized) {
		t.Error("Deserialized object does not match the original")
	}
}
