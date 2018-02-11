package main

import (
	"fmt"
)

func main() {
	ci1 := incrementor("Foo:")
	ci2 := incrementor("Bar:")
	p1 := puller(ci1)
	p2 := puller(ci2)

	fmt.Println("End Total", <-p1+<-p2)
}

func incrementor(s string) chan int {
	out := make(chan int)

	go func() {
		for i := 1; i < 20; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c chan int) chan int {
	out := make(chan int)

	go func() {
		var sum int
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()

	return out
}
