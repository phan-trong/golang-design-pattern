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

type calculator struct{}

func (c calculator) sumArea(shapes ...interface{}) float32 {
	var sum float32

	for _, shape := range shapes {
		switch shape.(type) {
		case circle:
			r := shape.(circle).radius
			sum += math.Pi * r * r
		case square:
			l := shape.(square).length
			sum += l * l
		}
	}

	return sum
}

func main() {
	circle := circle{radius: 1}
	square := square{length: 1}

	calculator := calculator{}

	fmt.Println(calculator.sumArea(circle, square))
}
