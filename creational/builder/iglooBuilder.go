// Concrete builder
package main

type iglooBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newIglooBuilder() *iglooBuilder {
    return &iglooBuilder{}
}

func (builder *iglooBuilder) setWindowType() {
    builder.windowType = "Snow Window"
}

func (builder *iglooBuilder) setDoorType() {
    builder.doorType = "Snow Door"
}

func (builder *iglooBuilder) setNumFloor() {
    builder.floor = 1
}

func (builder *iglooBuilder) getHouse() house {
    return house{
        doorType:   builder.doorType,
        windowType: builder.windowType,
        floor:      builder.floor,
    }
}