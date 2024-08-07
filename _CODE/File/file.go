package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err1 := os.Create("example.txt")

	if err1 != nil {
		fmt.Println("[-] Error while creating file: ", err1)
		return
	}

	defer file.Close()

	content := "Hello World!!!"
	_, err2 := io.WriteString(file, content)
	if err2 != nil {
		fmt.Println("[-] Error while creating file: ", err2)
		return
	}

	fmt.Println("[+] File Created.")

	// Reading file using buffer
	// file, err3 := os.Open("example.txt")

	// if err3 != nil {
	// 	fmt.Println("[-] Error while creating file: ", err3)
	// 	return
	// }
	// defer file.Close()

	// buffer := make([]byte, 1024)
	// for {
	// 	n, err4 := file.Read(buffer)
	// 	if err4 == io.EOF {
	// 		break
	// 	}
	// 	if err4 != nil {
	// 		fmt.Println("[-] Error while reading file: ", err4)
	// 		return
	// 	}
	// 	fmt.Println("Reading file using buffer: " + string(buffer[:n]))
	// }

	// Reading file using ioutil
	// content, err5 := os.ReadFile("example.txt")
	// if err5 != nil {
	// 	fmt.Println("[-] Error while reading file: ", err5)
	// 	return
	// }
	// fmt.Println("Reading file using os: " + string(content))
}
