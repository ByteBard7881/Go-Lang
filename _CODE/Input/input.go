package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Printf("NAME: %s", name)
}
