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

func (c *Country) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

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