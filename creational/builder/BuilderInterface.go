// Builder interface

package main

type BuilderInterface interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) BuilderInterface {
	
	switch builderType {
		case "normal":
			return &normalBuilder{}
		case "igloo":
			return 	&iglooBuilder{}
		default:
			return nil	
	}
}