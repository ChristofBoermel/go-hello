package main

import (
	"fmt"
)

func main() {
	name := ""
	age := 0
	height := 0.0


	fmt.Print("What is your name? ")

	fmt.Scan(&name)

	fmt.Print("What is your age? ")

	fmt.Scan(&age)

	fmt.Print("What is your height? ")

	fmt.Scan(&height)

	fmt.Printf("Hi %s, you are %d years old and %.2fm tall\n", name, age, height)
}
