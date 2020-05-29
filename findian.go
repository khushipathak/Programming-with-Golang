package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {

	var s string
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter a string: ")


	scanner.Scan()
	s = scanner.Text()
	
	s = strings.ToLower(s)
		
	if strings.HasPrefix(s, "i") {
		if strings.Contains(s, "a"){
			if strings.HasSuffix(s, "n"){
				fmt.Println("Found!")
			}  else {fmt.Println("Not Found!")}	
		}  else {fmt.Println("Not Found!")}	
	} else {fmt.Println("Not Found!")}	

}