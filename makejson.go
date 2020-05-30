package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	m := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter a name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Now enter an address: ")
	scanner.Scan()
	address := scanner.Text()

	m["address"] = address
	m["name"] = name

	barr, err := json.Marshal(m)

	fmt.Println("###################################")

	if err == nil {
		fmt.Println("JSON file: \n ", string(barr)) 
	} else {
		fmt.Println(err)
	}

	fmt.Println("###################################")

}
