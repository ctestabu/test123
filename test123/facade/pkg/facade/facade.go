package facade

import (
"fmt"
"history"
"notification"
"password"
"product"
"username"
"wallet"
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
