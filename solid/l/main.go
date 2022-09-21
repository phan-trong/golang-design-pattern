package main

import "fmt"

type person interface {
	getName() string
}

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	salary float64
}

type student struct {
	human
	marks float64
}

type printer struct{}

func (printer) info(p person) {
	fmt.Println("Name is: ", p.getName())
}

func main() {
	h := human{
		name: "human",
	}

	t := teacher{
		human: human{
			name: "teacher",
		},
		salary: 8000000,
	}

	s := student{
		human: human{
			name: "student",
		},
		marks: 9,
	}

	p := printer{}
	p.info(h)
	p.info(t)
	p.info(s)
}
