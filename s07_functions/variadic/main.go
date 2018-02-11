package main

import (
	"fmt"
)

func main() {
	n := average(43.0, 56.3, 32.6, 33.78)
	fmt.Println(n)

	args := []float64{12,23,45,56,23}
	n2 := average(args...)
	fmt.Println(n2)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)

	total := 0.0

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}