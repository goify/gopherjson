package gopherjson

type SerializableValue interface {
	Serialize() interface{}
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
