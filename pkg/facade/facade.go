package facade

import (
	"fmt"
	"reworked/pkg/history"
	"reworked/pkg/notification"
	"reworked/pkg/password"
	"reworked/pkg/product"
	"reworked/pkg/username"
	"reworked/pkg/wallet"
)

type CustomerFacade struct {
	username     *username.Username
	password     *password.Password
	wallet       *wallet.Wallet
	product      *product.Product
	notification *notification.Notification
	history      *history.History
}

//Adds value to the wallet and simple validation of pass and user
func (c *CustomerFacade) AddMoneyToWallet(usernameID string, password string, amount int) (err error) {
	fmt.Println("Adding money to your wallet")
	err = c.username.CheckUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.CheckPassword(password)
	if err != nil {
		return
	}
	c.wallet.CreditBalance(amount)
	c.notification.SendWalletNotification()
	c.history.NewEntry(usernameID, amount)
	return
}

//Substract value from wallet and simple validation of pass and user
func (c *CustomerFacade) SubtractMoneyFromWallet(usernameID string, password string, amount int) (err error) {
	fmt.Println("Substracting money from your wallet")
	err = c.username.CheckUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.CheckPassword(password)
	if err != nil {
		return
	}
	err = c.wallet.DeCreditBalance(amount)
	if err != nil {
		return
	}
	c.notification.SendWalletNotification()
	c.history.NewEntry(usernameID, amount)
	return
}

//Adds symbolic product to basket(slice) and simple validation of pass and user
func (c *CustomerFacade) AddProductToBasket(usernameID string, password string, prodID int) (err error) {
	fmt.Println("Adding product to basket")
	err = c.username.CheckUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.CheckPassword(password)
	if err != nil {
		return
	}
	err = c.product.AddNewProductToBasket(prodID) //
	if err != nil {
		return
	}
	c.notification.SendBasketNotification()
	c.history.NewProduct(usernameID, prodID)
	return
}

//Substract product from basket(slice) and simple validation of pass and user
func (c *CustomerFacade) SubtractProductToBasket(usernameID string, password string, prodID int) (err error) {
	fmt.Println("Adding product to basket")
	err = c.username.CheckUsername(usernameID)
	if err != nil {
		return
	}
	err = c.password.CheckPassword(password)
	if err != nil {
		return
	}
	err = c.product.RemoveProductFromBasket(prodID) //
	if err != nil {
		return
	}
	c.notification.SendBasketNotification()
	c.history.NewProduct(usernameID, prodID)
	return
}

//Вопрос. Правильно я понял, что эта функция инициализации и есть так
//называемая фабрика?
//Function to init CustomFacade structure
func NewCustomerFacade(usernameID string, usernamePassword string) *CustomerFacade {
	fmt.Println("Welcome to our brand new shop")
	CustomerFacade := &CustomerFacade{
		username:     username.New(usernameID),
		password:     password.New(usernamePassword),
		wallet:       wallet.New(),
		product:      product.New(),
		notification: &notification.Notification{},
		history:      &history.History{},
	}
	return CustomerFacade
}
