package main

import (
	"fmt"
)

func main() {
	s := greet("Cenk ", "Pekyaman")
	fmt.Println(s)
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}