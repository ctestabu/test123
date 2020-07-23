package username

import "fmt"

type Username struct {
	name string
}

//Username validation
func (u *Username) CheckUsername(newName string) (err error) {
	if u.name != newName {
		err = fmt.Errorf("Error in username")
	}
	fmt.Printf("11Welcome, %s\n", u.name)
	return
}

//Creates new username
func New(newName string) *Username {
	return &Username{
		name: newName,
	}
}
