package main

import "fmt"

func main() {
	number := 45
	name := "Maurya"

	var pointer1 *string = &name

	reference(&number)
	fmt.Println(pointer1)
	fmt.Println(*pointer1)
}

func reference(num *int) {
	fmt.Println(num)
	fmt.Println(*num)
}
