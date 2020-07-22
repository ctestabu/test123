package history

import "fmt"

type history struct {
}

func (h *history) newProduct(usernameID string, prodID int) {
	fmt.Printf("%s, added or removed item %d from basket", usernameID, prodID)
}

func (h *history) newEntry(usernameID string, amount int) {
	fmt.Printf("%s, added or removed money %d from basket", usernameID, amount)
}
