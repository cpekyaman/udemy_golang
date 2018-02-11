package main

import "fmt"

func wrapper() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {
	incr := wrapper()
	fmt.Println(incr())
	fmt.Println(incr())
}
