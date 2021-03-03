// Concrete decorator
package main

type tomatoTopping struct {
    pizza PizzaInterface
}

func (c *tomatoTopping) getPrice() int {
    pizzaPrice := c.pizza.getPrice()
    return pizzaPrice + 7
}
