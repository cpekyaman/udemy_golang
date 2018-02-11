package main

import (
	"fmt"
)

func main() {
	const start = 'A'
	const end = 'z'
	const size = end - start + 1

	var x [size]string

	fmt.Println(len(x))
	fmt.Printf("%T\n", x[0])
	fmt.Printf("%T\n", x)

	for i := start; i <= end; i++ {
		x[i - start] = string(i)
	}
	fmt.Println(x)
}