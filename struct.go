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

func clearName(u *user) {
	u.firstName = ""
}

func userain() {
	var u user
	fmt.Println("Enter Username")
	fmt.Scan(&u.firstName)
	fmt.Println("Enter Lasname")
	fmt.Scan(&u.lastName)
	fmt.Println("Enter Age")
	fmt.Scan(&u.age)
	userPrint(u)  // Here a new copy of u is created and is passed. If we want to work on changing the values of the variable u ( instead of creating a new value of it ), we need to pass a pointer instead of the variable
	clearName(&u) // Here, as we wanted to clear the firstname that user has given and stored in struct u.firstname, we pass the pointer to that struct. This prevents go to create a new copy of struct u but instead work on the original one
	userPrint(u)

}
