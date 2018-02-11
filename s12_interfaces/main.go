package main

import (
	"fmt"
	"math"
)

type Square struct {
	Side float64
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println("Area of shape is:", s.area())
}

func main() {
	s := Square{2.4}
	c := Circle{3.2}

	info(s)
	info(c)
}