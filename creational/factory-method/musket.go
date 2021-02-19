// Concrete product

package main

type musket struct {
    gun
}

func newMusket() GunInterface {
    return &musket{
        gun: gun{
            name:  "Musket gun",
            power: 1,
        },
    }
}