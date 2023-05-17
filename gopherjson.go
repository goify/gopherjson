package gopherjson

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
func Deserialize(serialized interface{}) (interface{}, error) {
	switch v := serialized.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})

		for key, val := range v {
			deserializedVal, err := Deserialize(val)

			if err != nil {
				return nil, err
			}

			result[key] = deserializedVal
		}

		return result, nil
	case []interface{}:
		result := make([]interface{}, len(v))

		for i, val := range v {
			deserializedVal, err := Deserialize(val)

			if err != nil {
				return nil, err
			}

			result[i] = deserializedVal
		}

		return result, nil
	case string:
		// Check for specific custom types and deserialize accordingly
		switch v {
		case "CustomDate":
			return CustomDate{}, nil // Adjust the deserialization logic according to your use case
		case "CustomRegex":
			return CustomRegex{}, nil // Adjust the deserialization logic according to your use case
		case "CustomFunction":
			return CustomFunction{}, nil // Adjust the deserialization logic according to your use case
		default:
			return v, nil
		}
	default:
		return serialized, nil
	}
}
