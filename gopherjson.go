package gopherjson

import (
	"fmt"
	"reflect"
)

type SerializableValue interface {
	Serialize() interface{}
	Deserialize(interface{}) error
}

// Serialize converts a value into its serializable form.
func Serialize(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case SerializableValue:
		return v.Serialize(), nil
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, val := range v {
			serializedVal, err := Serialize(val)
			if err != nil {
				return nil, err
			}
			result[key] = serializedVal
		}
		return result, nil
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, val := range v {
			serializedVal, err := Serialize(val)
			if err != nil {
				return nil, err
			}
			result[i] = serializedVal
		}
		return result, nil
	default:
		return value, nil
	}
}

// Deserialize converts a serialized value into its original form.
func Deserialize(serialized interface{}, value interface{}) error {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("value is not a valid pointer")
	}

	rv = rv.Elem()
	rt := rv.Type()

	switch v := serialized.(type) {
	case map[string]interface{}:
		switch rv.Kind() {
		case reflect.Struct:
			// Iterate over the struct fields and deserialize each field
			for i := 0; i < rv.NumField(); i++ {
				field := rv.Field(i)
				fieldType := rt.Field(i)
				fieldName := fieldType.Name

				if !field.CanSet() {
					continue
				}

				// Check if the field implements the SerializableValue interface
				if field.CanAddr() {
					if sv, ok := field.Addr().Interface().(SerializableValue); ok {
						if fieldValue, ok := v[fieldName]; ok {
							if err := sv.Deserialize(fieldValue); err != nil {
								return fmt.Errorf("error deserializing field %s: %v", fieldName, err)
							}
						}
						continue
					}
				}

				// Check if the field has a custom JSON tag
				tag := fieldType.Tag.Get("json")
				if tag != "" {
					fieldName = tag
				}

				// Deserialize the field value
				if fieldValue, ok := v[fieldName]; ok {
					if err := Deserialize(fieldValue, field.Addr().Interface()); err != nil {
						return fmt.Errorf("error deserializing field %s: %v", fieldName, err)
					}
				}
			}
		case reflect.Map:
			// Create a new map with the same type as the target value
			mapType := reflect.MapOf(rt.Key(), rt.Elem())
			mapValue := reflect.MakeMap(mapType)

			// Deserialize each key-value pair in the map
			for key, val := range v {
				keyValue := reflect.New(rt.Key()).Elem()
				valValue := reflect.New(rt.Elem()).Elem()

				// Deserialize the key and value
				if err := Deserialize(key, keyValue.Addr().Interface()); err != nil {
					return fmt.Errorf("error deserializing map key: %v", err)
				}
				if err := Deserialize(val, valValue.Addr().Interface()); err != nil {
					return fmt.Errorf("error deserializing map value: %v", err)
				}

				// Assign the key-value pair to the map
				mapValue.SetMapIndex(keyValue, valValue)
			}

			// Set the deserialized map as the value
			rv.Set(mapValue)
		default:
			return fmt.Errorf("unsupported type for deserialization: %v", rv.Kind())
		}
	case []interface{}:
		switch rv.Kind() {
		case reflect.Slice:
			// Create a new slice with the same type as the target value
			sliceType := reflect.SliceOf(rt.Elem())
			sliceValue := reflect.MakeSlice(sliceType, len(v), len(v))

			// Deserialize each element in the slice
			for i := 0; i < len(v); i++ {
				elemValue := reflect.New(rt.Elem()).Elem()

				// Deserialize the element
				if err := Deserialize(v[i], elemValue.Addr().Interface()); err != nil {
					return fmt.Errorf("error deserializing slice element: %v", err)
				}

				// Assign the element to the slice
				sliceValue.Index(i).Set(elemValue)
			}

			// Set the deserialized slice as the value
			rv.Set(sliceValue)
		default:
			return fmt.Errorf("unsupported type for deserialization: %v", rv.Kind())
		}
	case string:
		switch rv.Kind() {
		case reflect.Struct:
			if rv.Type() == reflect.TypeOf(CustomDate{}) {
				cd := CustomDate{}
				if err := cd.UnmarshalJSON([]byte(v)); err != nil {
					return fmt.Errorf("error deserializing CustomDate: %v", err)
				}
				rv.Set(reflect.ValueOf(cd))
			} else if rv.Type() == reflect.TypeOf(CustomRegex{}) {
				cr := CustomRegex{}
				if err := cr.UnmarshalJSON([]byte(v)); err != nil {
					return fmt.Errorf("error deserializing CustomRegex: %v", err)
				}
				rv.Set(reflect.ValueOf(cr))
			} else if rv.Type() == reflect.TypeOf(CustomFunction{}) {
				cf := CustomFunction{}
				if err := cf.UnmarshalJSON([]byte(v)); err != nil {
					return fmt.Errorf("error deserializing CustomFunction: %v", err)
				}
				rv.Set(reflect.ValueOf(cf))
			} else {
				return fmt.Errorf("unsupported type for deserialization: %v", rv.Kind())
			}
		default:
			return fmt.Errorf("unsupported type for deserialization: %v", rv.Kind())
		}
	default:
		// Set the serialized value directly if it's a primitive type
		rv.Set(reflect.ValueOf(serialized))
	}

	return nil
}
