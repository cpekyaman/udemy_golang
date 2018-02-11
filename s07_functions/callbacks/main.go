package main

import (
	"fmt"
)

func filter(nums []int, predicate func(int) bool) []int {
	var xs []int

	for _, n := range nums {
		if predicate(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func visit(nums []int, callback func(int)) {
	for _, n := range nums {
		callback(n)
	}
}

func main() {
	visit([]int{3,5,6,3,5}, func(n int) {
		fmt.Println(n)
	})

	xs := filter([]int{2,5,7,6,4,11,12}, func(n int) bool {
		return n % 2 == 0
	})
	fmt.Println(xs)
}