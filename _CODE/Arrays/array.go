package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Logan"
	name[1] = "Jake"
	name[2] = "Chris"

	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println("NAME: ", name)
	fmt.Println("LAST NAME: ", name[2])
	fmt.Println("LENGTH: ", len(numbers), numbers)
}
