package product

import "fmt"

type Product struct {
	id     int
	basket []int
}

//Adds product to basket(slice)
func (p *Product) AddNewProductToBasket(prodID int) (err error) {
	if prodID < 0 {
		err = fmt.Errorf("Not a valid id")
	}
	p.id = prodID
	p.basket = append(p.basket, prodID)
	return
}

//Removes product
func (p *Product) RemoveProductFromBasket(prodID int) (err error) {
	for _, i := range p.basket {
		if p.basket[i] == prodID {
			p.basket = append(p.basket[i:], p.basket[i+1:]...)
		} else {
			err = fmt.Errorf("Not a valid id")
		}
	}
	return
}

//Creates new product
func New() *Product {
	return &Product{
		id:     0,
		basket: []int{},
	}
}
