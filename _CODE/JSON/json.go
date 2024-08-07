package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person_object_1 := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println(person_object_1)

	// Encoding Object To Marshal (JSON Encoding)
	json_enoded_data, json_encode_error := json.Marshal(person_object_1)
	if json_encode_error != nil {
		fmt.Println("[-] Error while encoding object to json: ", json_encode_error)
	}
	// fmt.Println(json_enoded_data)
	fmt.Println(string(json_enoded_data))

	// Decoding Marshal (JSON Encoding) To Object
	var person_obecject_2 Person
	json_decode_error := json.Unmarshal(json_enoded_data, &person_obecject_2)
	if json_decode_error != nil {
		fmt.Println("[-] Error while decoding json to object: ", json_decode_error)
	}
	fmt.Println(person_obecject_2)
}
