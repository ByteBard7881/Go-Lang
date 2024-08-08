package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FakeAPI struct {
	User_Id   int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("CRUD Operations")
	url := "https://jsonplaceholder.typicode.com/todos/1"
	get(url)
}

func get(get_url string) {
	var get_fake_api FakeAPI

	get_response, get_error := http.Get(get_url)

	if get_error != nil {
		fmt.Println("[-] Error occured while performing get operation: ", get_error)
	}

	defer get_response.Body.Close()

	if get_response.StatusCode != http.StatusOK {
		fmt.Println("[-] Error recieving getting response: ", get_response.Status)
	}

	get_data, get_data_error := ioutil.ReadAll(get_response.Body)

	if get_data_error != nil {
		fmt.Println("[-] Error recieved in JSON Get Request: ", get_data_error)
		return
	}

	json_decoded_error := json.Unmarshal(get_data, &get_fake_api)

	if json_decoded_error != nil {
		fmt.Println("[-] Error while decoding json to object: ", json_decoded_error)
	}

	fmt.Println(get_fake_api)
	fmt.Println(string(get_data))
}
