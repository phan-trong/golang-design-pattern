package main

type SportBuilder struct {
	brand string
	color string
	door  int
}

func newSportBuilder() *SportBuilder {
	return &SportBuilder{}
}

func (b *SportBuilder) setBrand() {
	b.brand = "BMW"
}

func (b *SportBuilder) setColor() {
	b.color = "white"
}

func (b *SportBuilder) setNumDoor() {
	b.door = 2
}

func (b *SportBuilder) getCar() Car {
	return Car{
		brand: b.brand,
		color: b.color,
		door:  b.door,
	}
}
