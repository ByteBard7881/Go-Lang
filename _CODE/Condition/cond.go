package main

import "fmt"

func main() {
	x := 20

	if x < 5 {
		fmt.Println("X is smaller than 5")
	} else if x < 10 && x > 5 {
		fmt.Println("X is smaller than 10 and bigger than 5")
	} else {
		fmt.Println("X is bigger than 10")
	}
}
