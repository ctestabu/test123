package wallet

import (
	"fmt"
)

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
