// Abstract factory interface
package main

import "fmt"

type SportsFactoryInterface interface {
	makeShoe() ShoeInterface
	makeShirt() ShirtInterface
}

func getSportsFactory(brand string) (SportsFactoryInterface, error) {
	
	switch brand {
		case "adidas":
			return &adidas{}, nil
		case "nike":
			return &nike{}, nil
		default:
			return nil, fmt.Errorf("Wrong brand type passed")		
	}

}