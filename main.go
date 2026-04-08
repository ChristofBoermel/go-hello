package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("What is your name? ")

		name, _ := reader.ReadString('\n')


		name = strings.TrimSpace(name)

		fmt.Printf("Hello, %s\n", name)
		greeting := fmt.Sprintf("Hello, %s", name)
		//fmt.Print(greeting)
}
