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

func (c circle) area() {
	fmt.Println("Area of circle is: ", math.Pi*c.radius*c.radius)
}

func (s square) area() {
	fmt.Println("Area of square is: ", s.length*s.length)
}

func main() {
	circle := circle{radius: 1}
	circle.area()

	square := square{length: 1}
	square.area()
}
