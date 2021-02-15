// Director
package main

type director struct {
    builder BuilderInterface
}

func newDirector(builder BuilderInterface) *director {
    return &director{
        builder: builder,
    }
}

func (d *director) setBuilder(builder BuilderInterface) {
    d.builder = builder
}

func (d *director) buildHouse() house {
    d.builder.setDoorType()
    d.builder.setWindowType()
    d.builder.setNumFloor()
    return d.builder.getHouse()
}