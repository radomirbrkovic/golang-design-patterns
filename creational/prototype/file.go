// Concrete prototype

package main 
import "fmt"

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() NodeInterface {
	return &file{
		name: f.name + "_clone", 
	}
}