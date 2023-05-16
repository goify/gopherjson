package gopherjson

import (
	"fmt"
	"time"
)

// CustomDate represents a date that can be serialized and deserialized.
type CustomDate struct {
	time.Time
}

func (cd CustomDate) Serialize() interface{} {
	return cd.Format(time.RFC3339)
}

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	str := string(data)
	t, err := time.Parse(`"`+time.RFC3339+`"`, str)

	if err != nil {
		return fmt.Errorf("error unmarshaling CustomDate: %v", err)
	}

	cd.Time = t

	return nil
}
