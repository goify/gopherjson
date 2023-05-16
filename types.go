package gopherjson

import (
	"regexp"
	"time"
)

type SerializableValue interface {
	Serialize() interface{}
}

type CustomDate struct {
	time.Time
}

type CustomRegex struct {
	*regexp.Regexp
}
