package main

import "fmt"

const str string = "cenk"

const (
	PI = 3.14
	Language = "Go"
)

const (
	_ = iota
	b = iota * 10
	c = iota * 10
)

const (
	d = iota
)

func main() {
	const num int = 10

	fmt.Printf("str is %s\n", str)
	fmt.Printf("num is %d\n", num)

	fmt.Println(PI)
	fmt.Println(Language)

	fmt.Printf("b is %d and c is %d and d is %d", b, c, d)
}
