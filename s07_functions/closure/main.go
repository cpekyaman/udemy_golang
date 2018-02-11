package main

import (
	"fmt"
)

func main() {
	x := 42

	fmt.Printf("x is %d\n", x)

	{
		fmt.Printf("x is %d in inner scope \n", x)
		y := 55
		fmt.Printf("y is %d\n", y)
	}

	// fmt.println(y) does not compile
}