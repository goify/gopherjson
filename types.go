package gopherjson

type SerializableValue interface {
	Serialize() interface{}
}
