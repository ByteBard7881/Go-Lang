package main

import (
	"fmt"
	"net/url"
)

func main() {
	sample_url := "https://jsonplaceholder.typicode.com/"
	fmt.Printf("Type of sample url: %T\n", sample_url)

	parse_url, parse_err := url.Parse(sample_url)
	if parse_err != nil {
		fmt.Println("[-] Error while parsing url: ", parse_err)
	}
	fmt.Printf("Type of parsed url: %T", parse_url)
}
