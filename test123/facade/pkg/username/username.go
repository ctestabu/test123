package username

import (
	"fmt"
)

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
