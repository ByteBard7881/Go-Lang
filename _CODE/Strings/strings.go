package main

import (
	"fmt"
	"strings"
)

func main() {
	data1 := "apple,orange,banana"
	parts := strings.Split(data1, ",")
	fmt.Println(data1)
	fmt.Println(parts)

	data2 := "one two three two four three four"
	count := strings.Count(data2, "two")
	fmt.Println(count)

	data3 := "              Hello, Go!!!"
	trimmed := strings.TrimSpace(data3)
	fmt.Println(data3)
	fmt.Println(trimmed)

	name1 := "Aditya"
	name2 := "Singh"
	join := strings.Join([]string{name1, name2}, " ")
	fmt.Println(join)
}
