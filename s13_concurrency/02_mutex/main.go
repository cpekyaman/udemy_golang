package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
}

func increment(s string) {
	for i := 1; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
