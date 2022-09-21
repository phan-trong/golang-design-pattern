package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type square struct {
	length float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float32 {
	return s.length * s.length
}

func printArea(value float32) {
	fmt.Println("Area is: ", value)
}

func main() {
	c := circle{
		radius: 1,
	}

	printArea(c.area())

	s := square{
		length: 1,
	}

	printArea(s.area())
}
