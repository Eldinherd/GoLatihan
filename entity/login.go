package entity

import "fmt"

type Login struct {
	ID       int
	Username string
	Password string
}

func (l *Login) PrintDetail() {
	fmt.Println("ID:", l.ID)
	fmt.Println("Username : ", l.Username)
	fmt.Println("Password : ", l.Password)
}
