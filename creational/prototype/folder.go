// Concrete prototype

package main

import "fmt"

type folder struct {
	children []NodeInterface
	name string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
    for _, i := range f.children {
        i.print(indentation + indentation)
    }
}

func (f *folder) clone() NodeInterface {
	cloneFolder :=  &folder{name: f.name + "_clone"}
	var tmpChildren []NodeInterface

	for _, i := range f.children {
		copy := i.clone()
		tmpChildren = append(tmpChildren, copy)
	}

	cloneFolder.children = tmpChildren
	return cloneFolder
} 