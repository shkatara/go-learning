package main

import (
	"fmt"

	"example.com/struct/utils"
)

func main() {
	u := utils.NewUser("shubham", "katara", 21)

	a := utils.NewAdmin(u, "shubham.katara59@gmail.com", "redhat@123")

	err := utils.ValidateUser(u)
	if err != nil {
		fmt.Println(err)
	}
	utils.UserPrint(u)
	utils.AdminPrint(a)

	/*
		userPrint(u)  // Here a new copy of u is created and is passed. If we want to work on changing the values of the variable u ( instead of creating a new value of it ), we need to pass a pointer instead of the variable
		clearName(&u) // Here, as we wanted to clear the firstname that user has given and stored in struct u.firstname, we pass the pointer to that struct. This prevents go to create a new copy of struct u but instead work on the original one
		userPrint(u)
	*/
}
