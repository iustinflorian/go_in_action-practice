package main

import (
	"go_in_action-practice/chapter_5/entities"
)

type notifier interface {
	Notify()
}

func main() {
	a := entities.Admin{}
	a.New("iustin@gmail.com", "password")

	var n notifier
	n = &a

	n.Notify()
}
