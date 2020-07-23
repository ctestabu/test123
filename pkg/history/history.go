package history

import "fmt"

type History struct {
}

//simplified version of journal functional
func (h *History) NewProduct(usernameID string, prodID int) {
	fmt.Printf("%s, added or removed item %d from basket", usernameID, prodID)
}

//simplified version of journal functional
func (h *History) NewEntry(usernameID string, amount int) {
	fmt.Printf("%s, added or removed money %d from basket", usernameID, amount)
}
