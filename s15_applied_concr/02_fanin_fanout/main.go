package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4)

	c1 := sq(in)
	c2 := sq(in)
	c3 := sq(in)

	for n := range merge(c1, c2, c3) {
		fmt.Println(n)
	}
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func merge(ins ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(ins))

	for _, c := range ins {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
