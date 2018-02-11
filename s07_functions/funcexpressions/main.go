package main

import (
	"fmt"
)

func greeter() func() string {
	return func() string {
		return "Hello World II"
	}
}

func main() {
	greet := func() {
		fmt.Println("Hello World")
	}

	greet()

	greet2 := greeter()
	fmt.Println(greet2())
}