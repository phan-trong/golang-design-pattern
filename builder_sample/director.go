package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildCar() Car {
	d.builder.setBrand()
	d.builder.setColor()
	d.builder.setNumDoor()
	return d.builder.getCar()
}
