package main

import (
	"fmt"
)

func first3() {
	fmt.Println("K")
	fmt.Println("H")
	fmt.Println("U")
}

func last3() {
	fmt.Println("S")
	fmt.Println("H")
	fmt.Println("I")
}

func main() {

	fmt.Println("--------------------------------")
	fmt.Println("Two threads will attempt to print out two halves of my name KHUSHI")
	fmt.Println("- Enter 'yes' to see the race condition in action\n- Enter 'no' to view the intended execution of the two threads\n- Enter 'exit' to leave")
	fmt.Println("Tip: Run the race condition again and again for different outputs.")
	var resp string
	again := true

	for again{
		fmt.Scan(&resp)

		if resp == "yes" {
			go first3()
			go last3()
		} else if resp == "no" {
			first3()
			last3()
		} else if resp == "exit" {
			break;
		} else {
			fmt.Println("Invalid response. Please try again.")
			fmt.Print(">")
		}
	}

}

