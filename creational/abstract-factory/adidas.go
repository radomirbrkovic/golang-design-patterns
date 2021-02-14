// Concrete factory
package main

type adidas struct {
}

func (a *adidas) makeShoe() ShoeInterface {
	return &adidasShoe{
		shoe: shoe{
            logo: "adidas",
            size: 14,
        },
	}
}

func (a *adidas) makeShirt() ShirtInterface {
	return &adidasShirt {
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}