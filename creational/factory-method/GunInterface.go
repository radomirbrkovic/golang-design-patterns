// Product interface

package main

type GunInterface interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}