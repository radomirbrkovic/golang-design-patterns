// Concrete builder

package main 

type normalBuilder struct {
	windowType string
	doorType string
	floor int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (builder *normalBuilder) setWindowType() {
	builder.windowType = "Wooden Window"
}

func (builder *normalBuilder) setDoorType()  {
	builder.doorType = "Wooden Door"	
}

func (builder *normalBuilder) setNumFloor() {
	builder.floor = 2
}

func (builder *normalBuilder) getHouse() house {
	return house {
		doorType: builder.doorType,
		windowType: builder.windowType,
		floor: builder.floor,
	}
}