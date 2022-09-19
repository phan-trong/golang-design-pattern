package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	sportBuilder := getBuilder("sport")

	director := newDirector(normalBuilder)
	normalCar := director.buildCar()

	fmt.Printf("Normal Car Brand: %s\n", normalCar.brand)
	fmt.Printf("Normal Car Color: %s\n", normalCar.color)
	fmt.Printf("Normal Car Num Door: %d\n", normalCar.door)

	director.setBuilder(sportBuilder)
	sportCar := director.buildCar()

	fmt.Printf("Sport Car Brand: %s\n", sportCar.brand)
	fmt.Printf("Sport Car Color: %s\n", sportCar.color)
	fmt.Printf("Sport Car Num Door: %d\n", sportCar.door)
}
