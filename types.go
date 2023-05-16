package gopherjson

import "time"

type SerializableValue interface {
	Serialize() interface{}
}

type CustomDate struct {
	time.Time
}
