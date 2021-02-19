// Prototype Interface
package main

type NodeInterface interface {
	print(string)
	clone() NodeInterface
}