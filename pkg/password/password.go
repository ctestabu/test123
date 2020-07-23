package password

import "fmt"

type Password struct {
	password string
}

//Password check function
func (p *Password) CheckPassword(newPas string) (err error) {
	if p.password != newPas {
		err = fmt.Errorf("Error in password")
	}
	return
}

//Вопрос. Можно ли при разделении функционала по пакетам, создавать
//просто функцию New или надо указывать более подробную информацию?
//Creates new password
func New(newPas string) *Password {
	return &Password{
		password: newPas,
	}
}
