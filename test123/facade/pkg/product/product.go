package product

import (
	"fmt"
)

type product struct {
	id     int
	basket []int
}

func addProduct() *product {
	return &product{
		id:     0,
		basket: []int{},
	}
}

func (p *product) addNewProductToBasket(prodID int) (err error) {
	if prodID < 0 {
		err = fmt.Errorf("Not a valid id")
	}
	p.id = prodID
	p.basket = append(p.basket, prodID)
	return
}

func (p *product) removeProductFromBasket(prodID int) (err error) {
	for _, i := range p.basket {
		if p.basket[i] == prodID {
			p.basket = append(p.basket[i:], p.basket[i+1:]...) //!!!!
		} else {
			err = fmt.Errorf("Not a valid id")
		}
	}
	return
}
