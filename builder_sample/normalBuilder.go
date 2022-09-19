package main

type NormalBuilder struct {
	brand string
	color string
	door  int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setBrand() {
	b.brand = "Honda"
}

func (b *NormalBuilder) setColor() {
	b.color = "black"
}

func (b *NormalBuilder) setNumDoor() {
	b.door = 4
}

func (b *NormalBuilder) getCar() Car {
	return Car{
		brand: b.brand,
		color: b.color,
		door:  b.door,
	}
}
