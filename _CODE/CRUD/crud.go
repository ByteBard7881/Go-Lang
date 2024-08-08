package main

import (
	"encoding/json" // Package for encoding and decoding JSON
	"fmt"           // Package for formatted I/O
	"io/ioutil"     // Package for I/O utility functions
	"net/http"      // Package for HTTP client and server implementations
	"strings"       // Package for string manipulation
)

// Define a struct to represent the structure of data returned by the API
type FakeAPI struct {
	User_Id   int    `json:"userId"`    // JSON field mapping for user ID
	Id        int    `json:"id"`        // JSON field mapping for the unique identifier
	Title     string `json:"title"`     // JSON field mapping for the title
	Completed bool   `json:"completed"` // JSON field mapping for completion status
}

func main() {
	fmt.Println("CRUD Operations")

	// URLs for GET, POST, and PUT operations
	get_put_url := "https://jsonplaceholder.typicode.com/todos/1"
	post_url := "https://jsonplaceholder.typicode.com/todos"

	// Perform GET, POST, and PUT operations
	get(get_put_url)
	post(post_url)
	put(get_put_url)
}

// Function to perform a GET request
func get(get_url string) {
	var get_fake_api FakeAPI // Variable to hold the decoded JSON data

	get_response, get_error := http.Get(get_url) // Perform the GET request

	if get_error != nil { // Handle errors in GET request
		fmt.Println("[-] Error occured while performing get operation: ", get_error)
		return
	}

	defer get_response.Body.Close() // Ensure the response body is closed after function execution

	if get_response.StatusCode != http.StatusOK { // Check if the response status is OK
		fmt.Println("[-] Error recieving getting response: ", get_response.Status)
		return
	}

	get_data, _ := ioutil.ReadAll(get_response.Body) // Read the response body

	json_decoded_error := json.Unmarshal(get_data, &get_fake_api) // Decode the JSON response into the struct

	if json_decoded_error != nil { // Handle errors in JSON decoding
		fmt.Println("[-] Error while decoding json to object: ", json_decoded_error)
		return
	}

	// Print the response data and status
	fmt.Println("Get Response: ", string(get_data)+"   "+get_response.Status)
}

// Function to perform a POST request
func post(post_url string) {
	// Create a struct with the data to be sent in the POST request
	post_fake_data := FakeAPI{
		User_Id:   3456,
		Id:        854,
		Title:     "Lorem Nigga Ipsum",
		Completed: true,
	}

	// Convert the struct to JSON
	post_fake_api, post_data_error := json.Marshal(post_fake_data)

	if post_data_error != nil { // Handle errors in JSON encoding
		fmt.Println("[-] Error recieved in JSON Get Request: ", post_data_error)
		return
	}

	post_fake_api_reader := strings.NewReader(string(post_fake_api)) // Convert JSON to a readable format

	// Perform the POST request with JSON data
	post_response, post_error := http.Post(post_url, "application/json", post_fake_api_reader)

	if post_error != nil { // Handle errors in POST request
		fmt.Println("[-] Error while performing post method: ", post_error)
		return
	}

	defer post_response.Body.Close() // Ensure the response body is closed after function execution

	// Read and print the response body and status
	post_readable_format, _ := ioutil.ReadAll(post_response.Body)
	fmt.Println("Post Response: ", string(post_readable_format)+"   "+post_response.Status)
}

// Function to perform a PUT request
func put(put_url string) {
	// Create a struct with the data to be sent in the PUT request
	put_fake_api := FakeAPI{
		User_Id:   1,
		Id:        1,
		Title:     "Nigga",
		Completed: true,
	}

	// Convert the struct to JSON
	put_fake_json, put_fake_json_error := json.Marshal(put_fake_api)

	if put_fake_json_error != nil { // Handle errors in JSON encoding
		fmt.Println("[-] Error while encoding struct to json: ", put_fake_json_error)
		return
	}

	put_fake_json_reader := strings.NewReader(string(put_fake_json)) // Convert JSON to a readable format

	// Create a new PUT request
	put_json_request, put_json_request_error := http.NewRequest(http.MethodPut, put_url, put_fake_json_reader)

	if put_json_request_error != nil { // Handle errors in creating PUT request
		fmt.Println("[-] Error while creating put request: ", put_json_request_error)
		return
	}

	// Set the content type header to JSON
	put_json_request.Header.Set("Content-type", "application/json")

	client := http.Client{} // Create an HTTP client

	// Send the PUT request
	put_request, put_error := client.Do(put_json_request)

	if put_error != nil { // Handle errors in sending PUT request
		fmt.Println("[-] Error while sending put request: ", put_error)
		return
	}

	defer put_request.Body.Close() // Ensure the response body is closed after function execution

	// Read and print the response body and status
	put_json_request_bytes, _ := ioutil.ReadAll(put_request.Body)
	fmt.Println("Put Response: ", string(put_json_request_bytes)+"   "+put_request.Status)
}
