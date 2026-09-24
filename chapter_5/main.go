package main

import "fmt"

type Printer interface {
	Print()
}

type User struct {
	Name  string
	Email string // exported (public)
	id    int    // unexported (private)
}

type Admin struct {
	User
	Level string
}

func (u *User) Print() {
	fmt.Printf("User: %s | Email: %s\n", u.Name, u.Email)
}

func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	admin := &Admin{
		User: User{
			Name:  "Alex",
			Email: "alex@mail.com",
		},
		Level: "Super",
	}

	admin.ChangeEmail("alex@newdomain.com")

	var p Printer = admin
	p.Print()

	if concreteAdmin, ok := p.(*Admin); ok {
		fmt.Printf("Successfully asserted *Admin! Accessing Level: %s\n", concreteAdmin.Level)
	}

	inspectType(admin)
	inspectType(42)
	inspectType("Hello")
}

func inspectType(i any) {
	switch v := i.(type) {
	case *Admin:
		fmt.Printf("Type Switch: Received *Admin with Level '%s'\n", v.Level)
	case int:
		fmt.Printf("Type Switch: Received int value %d\n", v)
	case string:
		fmt.Printf("Type Switch: Received string '%s'\n", v)
	default:
		fmt.Println("Type Switch: Unknown type")
	}
}
