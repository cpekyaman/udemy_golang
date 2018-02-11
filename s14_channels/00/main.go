package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("c in")
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println("c out", <-c)
		}
	}()

	time.Sleep(time.Second)
}
