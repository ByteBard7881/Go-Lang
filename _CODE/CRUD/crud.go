package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type FakeAPI struct {
	User_Id   int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("CRUD Operations")
	post_url := "https://jsonplaceholder.typicode.com/todos"
	get_url := "https://jsonplaceholder.typicode.com/todos/1"
	get(get_url)
	post(post_url)
}

func get(get_url string) {
	var get_fake_api FakeAPI

	get_response, get_error := http.Get(get_url)

	if get_error != nil {
		fmt.Println("[-] Error occured while performing get operation: ", get_error)
		return
	}

	defer get_response.Body.Close()

	if get_response.StatusCode != http.StatusOK {
		fmt.Println("[-] Error recieving getting response: ", get_response.Status)
		return
	}

	get_data, _ := ioutil.ReadAll(get_response.Body)

	json_decoded_error := json.Unmarshal(get_data, &get_fake_api)

	if json_decoded_error != nil {
		fmt.Println("[-] Error while decoding json to object: ", json_decoded_error)
		return
	}

	fmt.Println(get_fake_api)
	fmt.Println(string(get_data))
}

func post(post_url string) {
	post_fake_data := FakeAPI{
		User_Id:   3456,
		Id:        854,
		Title:     "Lorem Nigga Ipsum",
		Completed: true,
	}

	post_fake_api, post_data_error := json.Marshal(post_fake_data)

	if post_data_error != nil {
		fmt.Println("[-] Error recieved in JSON Get Request: ", post_data_error)
		return
	}

	post_fake_api_reader := strings.NewReader(string(post_fake_api))

	post_response, post_error := http.Post(post_url, "application/json", post_fake_api_reader)

	if post_error != nil {
		fmt.Println("[-] Error while performing post method: ", post_error)
		return
	}

	defer post_response.Body.Close()

	post_readable_format, _ := ioutil.ReadAll(post_response.Body)
	fmt.Println(string(post_readable_format))
}
