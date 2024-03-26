package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	age       int
}

func userPrint(u user) {
	fmt.Println(u.age, u.lastName, u.firstName)
}

func struct_main() {
	var u user
	fmt.Println("Enter Username")
	fmt.Scan(&u.firstName)
	fmt.Println("Enter Lasname")
	fmt.Scan(&u.lastName)
	fmt.Println("Enter Age")
	fmt.Scan(&u.age)
	userPrint(u)
}
