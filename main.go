package main

import "fmt"

func buildGreeting(name string, age int, height float64) string {
	return fmt.Sprintf("Hi %s, you are %d years old and %.2fm tall", name, age, height)
}

func greet(name string, age int, height float64) {
	message := buildGreeting(name, age, height)
	fmt.Println(message)
}

func getUserInfo() (string, int, float64) {
	name := ""
	age := 0
	height := 0.0

	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	fmt.Print("How old are you? ")
	fmt.Scan(&age)

	fmt.Print("How tall are you? ")
	fmt.Scan(&height)

	return name, age, height
}

func main() {
	name, age, height := getUserInfo()
	greet(name, age, height)
}
