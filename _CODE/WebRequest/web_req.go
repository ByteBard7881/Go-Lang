package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	get_response, get_err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if get_err != nil {
		fmt.Println("[-] Error recieved in JSON Get Request: ", get_err)
		return
	}
	defer get_response.Body.Close()

	// Read Response
	data, data_err := ioutil.ReadAll(get_response.Body)
	if data_err != nil {
		fmt.Println("[-] Error recieved in JSON Get Request: ", data_err)
		return
	}
	fmt.Println("HTTP Get JSON Response:\n" + string(data))
}
