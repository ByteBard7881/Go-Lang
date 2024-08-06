package main

import "fmt"

func main() {
	maps := make(map[string]int)

	maps["k1"] = 7
	maps["k2"] = 13

	fmt.Println(maps)

	variable := maps["k1"]
	fmt.Println(variable)
	fmt.Println(len(maps))

	delete(maps, "k2")
	fmt.Println(maps)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
