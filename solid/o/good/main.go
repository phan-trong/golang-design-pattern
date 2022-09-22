package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

type circle struct {
	radius float32
}

type square struct {
	length float32
}

type calculator struct{}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float32 {
	return s.length * s.length
}

func (c calculator) sumArea(shapes ...shape) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func main() {
	circle := circle{radius: 1}
	square := square{length: 1}

	calculator := calculator{}

	fmt.Println(calculator.sumArea(circle, square))
}
