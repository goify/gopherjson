# GopherJSON

`gopherjson` is a powerful and flexible Go package that provides advanced serialization and deserialization capabilities for working with JSON data.

## Features

- Serialize custom types: Convert custom types into their serializable form for JSON representation.
- Deserialize custom types: Convert serialized JSON values back into their original custom types.
- Support for `CustomDate`: Serialize and deserialize `time.Time` values with custom date format.
- Support for `CustomRegex`: Serialize and deserialize regular expressions.
- Support for `CustomFunction`: Serialize and deserialize functions.

## Installation

To use `gopherjson` in your Go project, you need to install it using the `go get` command:

```bash
go get github.com/iamando/gopherjson
```

## Usage

Here's a simple example demonstrating how to use `gopherjson`:

```go
package main

import (
 "fmt"
 "time"

 "github.com/iamando/gopherjson"
)

type Person struct {
 Name      string
 BirthDate gopherjson.CustomDate
}

func main() {
 // Create an instance of the custom type
 p := Person{
  Name:      "John Doe",
  BirthDate: gopherjson.CustomDate{time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)},
 }

 // Serialize the custom type
 serialized, err := gopherjson.Serialize(p)
 if err != nil {
  fmt.Println("Serialization error:", err)
  return
 }

 // Deserialize the serialized data
 var deserialized Person
 err := gopherjson.Deserialize(serialized, &deserialized)
 if err != nil {
  fmt.Println("Deserialization error:", err)
  return
 }

 // Convert the deserialized value back to the original type
 converted, ok := deserialized.(Person)
 if !ok {
  fmt.Println("Failed to convert to Person type")
  return
 }

 // Print the results
 fmt.Println("Original Person:", p)
 fmt.Println("Serialized Data:", serialized)
 fmt.Println("Deserialized Person:", converted)
}
```

## Support

GopherJSON is an MIT-licensed open source project. It can grow thanks to the sponsors and support.

## License

GopherJSON is [MIT licensed](LICENSE).
