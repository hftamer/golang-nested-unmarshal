package main

import (
	"encoding/json"
	"fmt"
)

type (
	Country struct {
		Name string `json:"name"`
	}

	Person struct {
		Id      int     `json:"id"`
		Country Country `json:"country"`
	}
)

// UnmarshalJSON note that this is a member function on the object
// we wish to deserialize
func (c *Country) UnmarshalJSON(data []byte) error {
	// Normally, unmarshaling this struct would require it to be defined
	// as a JSON object

	// In our case, the country is defined as a string on the payload
	// so we have to deserialize it as a string
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	// If deserialization is successful, assign it to the object
	c.Name = name

	return nil
}

func main() {
	data := `{
		"id": 1,
		"country": "US"
	}`

	fmt.Println("Unmarshalling string data:\n", data)

	person := Person{}
	if err := json.Unmarshal([]byte(data), &person); err != nil {
		panic("unable to unmarshal data")
	}

	fmt.Println("Deserialized Person:\n", person)
}