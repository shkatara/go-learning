package main

import (
	"fmt"

	"example.com/struct/utils"
)

func main() {
	u := utils.User{}
	fmt.Println("Enter Username")
	fmt.Scanln(&u.FirstName)
	fmt.Println("Enter Lasname")
	fmt.Scanln(&u.LastName)
	fmt.Println("Enter Age")
	fmt.Scanln(&u.Age)
	err := utils.ValidateUser(u)
	if err != nil {
		fmt.Println(err)
	}

	/*
		userPrint(u)  // Here a new copy of u is created and is passed. If we want to work on changing the values of the variable u ( instead of creating a new value of it ), we need to pass a pointer instead of the variable
		clearName(&u) // Here, as we wanted to clear the firstname that user has given and stored in struct u.firstname, we pass the pointer to that struct. This prevents go to create a new copy of struct u but instead work on the original one
		userPrint(u)
	*/
}
