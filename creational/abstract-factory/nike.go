// Concrete factory
package main

type nike struct {
}

func (n *nike) makeShoe() ShoeInterface {
	return &nikeShoe {
		shoe: shoe{
            logo: "adidas",
            size: 14,
        },
	}
}

func (n *nike) makeShirt() ShirtInterface  {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}