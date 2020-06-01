package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct {
	food     string
	movement string
	sound    string
}

type Bird struct {
	food     string
	movement string
	sound    string
}

type Snake struct {
	food     string
	movement string
	sound    string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.movement)
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.movement)
}

func (b Bird) Speak() {
	fmt.Println(b.sound)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.movement)
}

func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func MakeAnimal(a Animal) bool {
	fmt.Print("Created it!\n>")
	return true
}

func GiveInfo(a Animal, info string) {

	cow, ok := a.(Cow)
	if ok {
		switch info {
		case "eat":
			cow.Eat()
			fmt.Print(">")
		case "move":
			cow.Move()
			fmt.Print(">")
		case "speak":
			cow.Speak()
			fmt.Print(">")
		default:
			fmt.Println("Invalid query, please try again.")
			fmt.Print(">")
		}
	}

	bird, ok := a.(Bird)
	if ok {
		switch info {
		case "eat":
			bird.Eat()
			fmt.Print(">")
		case "move":
			bird.Move()
			fmt.Print(">")
		case "speak":
			bird.Speak()
			fmt.Print(">")
		default:
			fmt.Println("Invalid query, please try again.")
			fmt.Print(">")
		}
	}

	snake, ok := a.(Snake)
	if ok {
		switch info {
		case "eat":
			snake.Eat()
			fmt.Print(">")
		case "move":
			snake.Move()
			fmt.Print(">")
		case "speak":
			snake.Speak()
			fmt.Print(">")
		default:
			fmt.Println("Invalid query, please try again.")
			fmt.Print(">")
		}
	}

}

func main() {

	fmt.Print("Hello!\n- To make a new animal, type 'newanimal [animal_name] [animal_type]'\n- To view details about an animal already created, type 'query [animal_type] [info]'\nType 'exit' to quit\n>")

	scanner := bufio.NewScanner(os.Stdin)

	cow_bool := false
	bird_bool := false
	snake_bool := false

	for scanner.Scan() {

		req := scanner.Text()

		if req == "exit" {
			break
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Invalid input, please rerun the program.")
			}
		}()

		req_sli := strings.Split(req, " ")

		req_name := req_sli[0]

		switch req_name {

		case "newanimal":

			animal_type := req_sli[2]

			switch animal_type {
			case "cow":
				cow := Cow{food: "grass", sound: "moo", movement: "walk"}
				cow_bool = MakeAnimal(cow)
			case "bird":
				bird := Bird{food: "worms", sound: "peep", movement: "fly"}
				bird_bool = MakeAnimal(bird)
			case "snake":
				snake := Snake{food: "mice", sound: "hsss", movement: "slither"}
				snake_bool = MakeAnimal(snake)
			default:
				fmt.Print("Invalid animal\n>")
			}

		case "query":

			animal_type := req_sli[1]
			info := req_sli[2]

			switch animal_type {
			case "cow":

				if cow_bool {
					cow := Cow{food: "grass", sound: "moo", movement: "walk"}
					GiveInfo(cow, info)
				} else {
					fmt.Print("No cow created yet. Please try again\n>")
				}

			case "bird":

				if bird_bool {
					bird := Bird{food: "worms", sound: "peep", movement: "fly"}
					GiveInfo(bird, info)
				} else {
					fmt.Print("No bird created yet. Please try again\n>")
				}

			case "snake":

				if snake_bool {
					snake := Snake{food: "mice", sound: "hsss", movement: "slither"}
					GiveInfo(snake, info)
				} else {
					fmt.Print("No snake created yet. Please try again\n>")
				}

			}
		}

	}

}
