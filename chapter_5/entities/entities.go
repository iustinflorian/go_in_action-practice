package entities

import "fmt"

type user struct {
	Email    string
	password string
}

type Admin struct {
	user
	Role string
}

func (a *Admin) New(email string, password string) {
	a.user = user{
		Email:    email,
		password: password,
	}
	a.Role = "default"
}

func (u *user) Notify() {
	fmt.Printf("User with email %s has been notified", u.Email)
}
