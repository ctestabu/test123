package password

import (
	"fmt"
)

type password struct {
	password string
}

func newPassword(newPas string) *password {
	return &password{
		password: newPas,
	}
}

func (p *password) checkPassword(newPas string) (err error) {
	if p.password != newPas {
		err = fmt.Errorf("Error in password")
	}
	return
}
