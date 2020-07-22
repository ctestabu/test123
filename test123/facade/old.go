package main

import (
	"fmt"
	"log"
)

type customerFacade struct {
	username     *username
	password     *password
	wallet       *wallet
	product      *product
	notification *notification
	history      *history
}

func newCustomerFacade(usernameID string, usernamePassword string) *customerFacade {
	fmt.Println("Welcome to our brand new shop")
	customerFacade := &customerFacade{
		username:     newUsername(usernameID),
		password:     newPassword(usernamePassword),
		wallet:       newWallet(),
		product:      addProduct(),
		notification: &notification{},
		history:      &history{},
	}
	return customerFacade
}

//customerfacade functions

//add money

func (c *customerFacade) addMoneyToWallet(usernameID string, password string, amount int) (err error) {
	fmt.Println("Adding money to your wallet")
	err = c.username.checkUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.checkPassword(password)
	if err != nil {
		return
	}
	c.wallet.creditBalance(amount)
	c.notification.sendWalletNotification()
	c.history.newEntry(usernameID, amount)
	return
}

func (c *customerFacade) subtractMoneyFromWallet(usernameID string, password string, amount int) (err error) {
	fmt.Println("Substracting money from your wallet")
	err = c.username.checkUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.checkPassword(password)
	if err != nil {
		return
	}
	err = c.wallet.deCreditBalance(amount)
	if err != nil {
		return
	}
	c.notification.sendWalletNotification()
	c.history.newEntry(usernameID, amount)
	return
}

func (c *customerFacade) addProductToBasket(usernameID string, password string, prodID int) (err error) {
	fmt.Println("Adding product to basket")
	err = c.username.checkUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.checkPassword(password)
	if err != nil {
		return
	}
	err = c.product.addNewProductToBasket(prodID) //
	if err != nil {
		return
	}
	c.notification.sendBasketNotification()
	c.history.newProduct(usernameID, prodID)
	return
}

func (c *customerFacade) subtractProductToBasket(usernameID string, password string, prodID int) (err error) {
	fmt.Println("Adding product to basket")
	err = c.username.checkUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.checkPassword(password)
	if err != nil {
		return
	}
	err = c.product.removeProductFromBasket(prodID) //
	if err != nil {
		return
	}
	c.notification.sendBasketNotification()
	c.history.newProduct(usernameID, prodID)
	return
}

//username.go

type username struct {
	name string
}

func newUsername(newName string) *username {
	return &username{
		name: newName,
	}
}

func (u *username) checkUsername(newName string) (err error) {
	if u.name != newName {
		err = fmt.Errorf("Error in username")
	}
	fmt.Printf("11Welcome, %s\n", u.name)
	return
}

//password.go

type password struct {
	password string
}

func newPassword(newPas string) *password {
	return &password{
		password: newPas,
	}
}

func (p *password) checkPassword(newPas string) (err error) {
	if p.password != newPas {
		err = fmt.Errorf("Error in password")
	}
	return
}

//wallet.go

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Added balance to wallet")
	return
}

func (w *wallet) deCreditBalance(amount int) (err error) {
	if w.balance < amount {
		err = fmt.Errorf("Not enough money")
	}
	fmt.Println("- babki")
	w.balance = w.balance - amount
	return
}

// product.go
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

//notification

type notification struct {
}

func (n *notification) sendWalletNotification() {
	fmt.Println("Pay attention to operations with your wallet")
}

func (n *notification) sendBasketNotification() {
	fmt.Println("Pay attention to operations with your basket")
}

//history

type history struct {
}

func (h *history) newProduct(usernameID string, prodID int) {
	fmt.Printf("%s, added or removed item %d from basket", usernameID, prodID)
}

func (h *history) newEntry(usernameID string, amount int) {
	fmt.Printf("%s, added or removed money %d from basket", usernameID, amount)
}

func main() {
	fmt.Println("-----------")
	example := newCustomerFacade("Johnt", "passtest")
	fmt.Println("---------------------")
	err := example.addMoneyToWallet("Johnt", "passtest", 150)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = example.subtractMoneyFromWallet("test", "passtest", 50)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
