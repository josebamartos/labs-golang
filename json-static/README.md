# Static JSON

Examples on this directory shows how to marshal data to a JSON string and how to unmarshal it back.

Marshaling and unmarshaling is done statically using [structs](https://golang.org/ref/spec#Struct_types) and the tools of [encoding/json](https://golang.org/pkg/encoding/json/) package.

Sample data is about the [brightest stars in the sky](https://en.wikipedia.org/wiki/List_of_brightest_stars).

| Directory                                        | Description                                                                                                                                                                                                                                        |
| ------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [json-slice-of-structs](json-array-of-structs)   | Creates a slice of "Star" structs. Marshals the slice to a JSON string, unmarshals the string back to a new slice, and iterates over the items of the slice printing the values of each one.                                                       |
| [json-struct-with-slice](json-struct-with-array) | Creates the "Stars" struct containing a slice of "Star" structs. Marshals the "Stars" struct to a JSON string, unmarshals the string back to a new "Stars" struct, and iterates over the items of the inner slice printing the values of each one. |
| [json-nested-structs](json-nested-structs)       | Creates the "Stars" struct containing a "Star" struct for each item. Marshals the "Stars" struct to a JSON string, unmarshals the string back to a new "Stars" struct, and prints the values of each inner struct.                                 |
