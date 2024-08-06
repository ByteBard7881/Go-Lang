package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()
	fmt.Println("Current Time: ", time)

	format_time := time.Format("02-01-2006, Monday, 15:04:05.00")
	fmt.Println(format_time)
}
