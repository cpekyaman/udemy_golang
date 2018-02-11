package main

import (
	"fmt"
)

type Printable interface {
	PrintSelf()
}

type MyInt int

type Person struct {
	First string
	Last string
	Age int
}

type DoubleZero struct {
	Person
	First string
	Code string
}

func (p Person) PrintSelf() {
	fmt.Println("First:", p.First, "Last:", p.Last)
}
func (p Person) FullName() string {
	return p.First + " " + p.Last
}

func main() {
	jb := DoubleZero{
		Person: Person{
			First : "James",
			Last : "Bond",
			Age : 30,
		},
		First : "Double Zero Seven",
		Code : "007",
	}

	jb.PrintSelf()
	fmt.Println(jb.FullName())

	var mi MyInt = 10
	var i int = 10

	fmt.Println(int(mi) + i)
}