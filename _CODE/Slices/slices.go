package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)

	numbers = append(numbers, 6, 7, 8, 9, 10)
	fmt.Println(numbers)

	name := []string{}
	fmt.Println(name)
}
