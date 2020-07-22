package notification

import "fmt"

type notification struct {
}

func (n *notification) sendWalletNotification() {
	fmt.Println("Pay attention to operations with your wallet")
}

func (n *notification) sendBasketNotification() {
	fmt.Println("Pay attention to operations with your basket")
}
