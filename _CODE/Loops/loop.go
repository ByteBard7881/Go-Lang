package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	names := []string{}

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter name at %d index> ", i)
		read := bufio.NewReader(os.Stdin)
		data, _ := read.ReadString('\n')

		names = append(names, data)
	}

	for i := 0; i < len(names); i++ {
		fmt.Printf("Names in array using loops: %s", names[i])
	}

	counter := 0
	for {
		if counter == 3 {
			counter++
			continue
		} else if counter == 10 {
			break
		}
		fmt.Println(counter)
		counter++
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Value at index %d is> %d\n", index, value)
	}
}
