package notification

import "fmt"

type Notification struct {
}

//simplified version of notification functional
func (n *Notification) SendWalletNotification() {
	fmt.Println("Pay attention to operations with your wallet")
}


//simplified version of notification functional
func (n *Notification) SendBasketNotification() {
	fmt.Println("Pay attention to operations with your basket")
}
