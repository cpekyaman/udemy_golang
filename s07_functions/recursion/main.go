package main

import (
	"fmt"
)

func accumulate(acc int32, n int32) int32 {
	if n == 0 {
		return acc
	}
	return accumulate(n * acc, n - 1)
}

func fact(n int32) int32 {
	return accumulate(1, n)
}

func main() {
	fmt.Println(fact(6))
}