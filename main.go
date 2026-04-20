package main 

import (
	"fmt"
	"strings"	
)

type Person struct {
	Name string
	Age int	
	Height float64
	Email string
	Gender string
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email address: %s", email)
	}
	return nil
}

func getUserInfo() (Person, error) {
	p := Person{}

	fmt.Println("What is your name?: ")
	_, err := fmt.Scan(&p.Name)
	if err != nil	{
			return Person{}, err
	}

	fmt.Println("How old are you?: ")
	_, err = fmt.Scan(&p.Age)
	if err != nil {
			return Person{}, err
	}

	fmt.Println("How tall are you?: ")
	_, err = fmt.Scan(&p.Height)
	if err != nil {
			return Person{}, err
	}

	fmt.Println("What is your gender, m, f or nb?: ")
	_, err = fmt.Scan(&p.Gender)
	if err != nil {
		return Person{}, err
	}

	fmt.Println("Type in your email!: ")
	_, err = fmt.Scan(&p.Email)
	if err != nil {
		return Person{}, err
	}
	err = validateEmail(p.Email)
	if err != nil {
		return Person{}, err
	}
	return p, nil
}

func greet(p Person) {
	fmt.Printf("Hi %s, you are %d years old and %.2fm tall, you are %s and your email is %s\n", p.Name, p.Age, p.Height, p.Gender, p.Email)
}

func main() {
	p, err := getUserInfo()
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}
	greet(p)
}
