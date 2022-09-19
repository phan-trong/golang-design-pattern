package main

type IBuilder interface {
	setBrand()
	setColor()
	setNumDoor()
	getCar() Car
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "sport" {
		return newSportBuilder()
	}

	return nil
}
