package main

import "fmt"

// Two separated interface

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64

	for _, s := range shapes {
		sum += s.area()
	}

	return sum
}

func areaVolumeSum(shapes ...object) float64 {
	var sum float64

	for _, s := range shapes {
		sum += s.area() + s.volume()
	}

	return sum
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type cube struct {
	length float64
}

func (c cube) area() float64 {
	return 6.0 * c.length * c.length
}

func (c cube) volume() float64 {
	return c.length * c.length * c.length
}

func main() {
	s := square{length: 1}
	c := cube{length: 2}

	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(c))
}
