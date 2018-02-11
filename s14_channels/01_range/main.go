package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("c in")
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println("n out", n)
	}
}
