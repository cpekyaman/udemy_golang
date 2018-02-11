package main

import (
	"fmt"
)

func multidim() {
	records := make([][]string, 5, 10)

	one := make([]string, 3)
	one[0] = "cenk"
	one[1] = "pekyaman"
	one[2] = "architect"

	records[0] = one

	fmt.Println(records)
	fmt.Println(records[0])
	fmt.Println(records[0][2])
}

func basics() {
	slice := make([]int, 5, 10)

	fmt.Println(slice)
	for i := 0; i < 20; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(slice[5:11])
}

func main() {
	basics()
	multidim()
}