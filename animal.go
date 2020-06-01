package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food     string
	sound    string
	movement string
}

func (a *animal) Eat() string {
	return a.food
}

func (a *animal) Move() string {
	return a.movement
}

func (a *animal) Speak() string {
	return a.sound
}

func main() {

	cow := animal{food: "grass", sound: "moo", movement: "walk"}
	bird := animal{food: "worms", sound: "peep", movement: "fly"}
	snake := animal{food: "mice", sound: "hsss", movement: "slither"}

	fmt.Println("Enter your request in the form of '[animal] [detail]', type 'exit' to leave\nFor eg. cow speak\n>")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		req := scanner.Text()

		if req == "exit" {
			break
		}

		req_sli := strings.Split(req, " ")

		name := req_sli[0]
		info := req_sli[1]

		switch name {

		case "cow":
			switch info {
			case "eat":
				fmt.Println(cow.Eat())
				fmt.Print(">")
			case "move":
				fmt.Println(cow.Move())
				fmt.Print(">")
			case "speak":
				fmt.Println(cow.Speak())
				fmt.Print(">")
			}
		case "bird":
			switch info {
			case "eat":
				fmt.Println(bird.Eat())
				fmt.Print(">")
			case "move":
				fmt.Println(bird.Move())
				fmt.Print(">")
			case "speak":
				fmt.Println(bird.Speak())
				fmt.Print(">")
			}
		case "snake":
			switch info {
			case "eat":
				fmt.Println(snake.Eat())
				fmt.Print(">")
			case "move":
				fmt.Println(snake.Move())
				fmt.Print(">")
			case "speak":
				fmt.Println(snake.Speak())
				fmt.Print(">")
			}
		default:
			fmt.Println("ERROR. INVALID INPUT")
			fmt.Print(">")

		}
	}

}
