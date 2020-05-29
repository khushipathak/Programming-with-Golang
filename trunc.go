package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello there!")

	var num float32
	fmt.Println("Enter a floating number: ")
	fmt.Scan(&num)

	fmt.Println("The truncated integer is: ")
	fmt.Println(int64(num))

}
