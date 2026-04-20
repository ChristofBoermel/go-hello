package main

import "fmt"

func buildGreeting(name string, age int, height float64) string {
	return fmt.Sprintf("Hi %s, you are %d years old and %.2fm tall", name, age, height)
}

func greet(name string, age int, height float64) {
	message := buildGreeting(name, age, height)
	fmt.Println(message)
}

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("invalid age: %d, must be between 0 and 150", age)
	}
	return nil
}

func getUserInfo() (string, int, float64, error) {
	name := ""
	age := 0
	height := 0.0

	fmt.Print("What is your name? ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return "", 0, 0.0, err
	}

	fmt.Print("How old are you? ")
	fmt.Scan(&age)
	err = validateAge(age)	
		if err != nil {
		return "", 0, 0.0, err
	}

	fmt.Print("How tall are you? ")
	_, err = fmt.Scan(&height)
		if err != nil {
		return "", 0, 0.0, err
	}

	return name, age, height, nil
}

func main() {
	name, age, height, err := getUserInfo()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
		return
	}
	greet(name, age, height)
}
