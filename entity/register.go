package entity

import (
	"fmt"
)

type Register struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
}

func (r *Register) PrintDetail() {
	fmt.Println("Username : ", r.Username)
	fmt.Println("FirstName : ", r.FirstName)
	fmt.Println("LastName : ", r.LastName)
	fmt.Println("Password : ", r.Password)
}
