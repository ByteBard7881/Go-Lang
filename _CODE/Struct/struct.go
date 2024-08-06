package main

import "fmt"

type Person struct {
	FName string
	LNAME string
	AGE   int
}

func main() {
	var person Person
	fmt.Println(person)

	person.FName = "Logan"
	person.LNAME = "Jackman"
	person.AGE = 55

	fmt.Println(person)

	person1 := Person{
		FName: "Caleb",
		LNAME: "Thomas",
		AGE:   34,
	}

	fmt.Println(person1)
}
