package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var param string

	fmt.Printf("Enter a message> ")
	read := bufio.NewReader(os.Stdin)
	param, _ = read.ReadString('\n')

	message(param)
}

func message(msg string) {
	fmt.Printf("%s", msg)
}
