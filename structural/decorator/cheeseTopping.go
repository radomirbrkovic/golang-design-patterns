// Concrete decorator
package main

type cheeseTopping struct {
    pizza PizzaInterface
}

func (c *cheeseTopping) getPrice() int {
    pizzaPrice := c.pizza.getPrice()
    return pizzaPrice + 10
}
