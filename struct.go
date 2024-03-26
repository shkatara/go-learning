package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	age       int
}

func (u user) userPrint() {
	fmt.Println(u.age, u.lastName, u.firstName)
}

func user_main() {
	var u user
	fmt.Println("Enter Username")
	fmt.Scan(&u.firstName)
	fmt.Println("Enter Lasname")
	fmt.Scan(&u.lastName)
	fmt.Println("Enter Age")
	fmt.Scan(&u.age)
	u.userPrint()
}
