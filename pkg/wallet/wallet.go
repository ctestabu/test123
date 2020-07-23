package wallet

import "fmt"

type Wallet struct {
	balance int
}

//Adds value to wallet balance
func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Added balance to wallet")
	return
}

//Removes value to wallet balance
func (w *Wallet) DeCreditBalance(amount int) (err error) {
	if w.balance < amount {
		err = fmt.Errorf("Not enough money")
	}
	fmt.Println("- babki")
	w.balance = w.balance - amount
	return
}

//Creates new wallet
func New() *Wallet {
	return &Wallet{
		balance: 0,
	}
}
