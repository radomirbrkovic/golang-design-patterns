// Concrete receiver

package main

import "fmt"

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turing tv on")
}

func (t *tv) off()  {
	t.isRunning = false
	fmt.Println("Turing tv off")
}