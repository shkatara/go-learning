package main

import (
	"errors"
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

func validateUser(u user) error {
	if u.firstName == "" || u.lastName == "" || u.age < 20 {
		return errors.New("First and last Name can not be empty and age should be more than 20")
	}
	return nil
}

func main() {
	var u user
	fmt.Println("Enter Username")
	fmt.Scanln(&u.firstName)
	fmt.Println("Enter Lasname")
	fmt.Scanln(&u.lastName)
	fmt.Println("Enter Age")
	fmt.Scanln(&u.age)
	err := validateUser(u)
	if err != nil {
		fmt.Println(err)
	}

	/*
		userPrint(u)  // Here a new copy of u is created and is passed. If we want to work on changing the values of the variable u ( instead of creating a new value of it ), we need to pass a pointer instead of the variable
		clearName(&u) // Here, as we wanted to clear the firstname that user has given and stored in struct u.firstname, we pass the pointer to that struct. This prevents go to create a new copy of struct u but instead work on the original one
		userPrint(u)
	*/
}
