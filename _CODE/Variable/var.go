package main

import "fmt"

func main() {
	var name string = "Aditya"
	var last_name = "Singh"
	var money int = 650000
	var pi float64 = 3.14
	var decided bool = true
	const hero string = "logan"
	fmt.Println(name + " " + last_name)
	fmt.Println(money)
	fmt.Println(pi)
	fmt.Println(decided)
	fmt.Println(hero)

	person := 123
	fmt.Println(person)

	var Public = "Gonna be used by other files"
	fmt.Println(Public)
}
