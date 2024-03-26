package utils

import (
	"errors"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

type Admin struct {
	Email    string
	Password string
	User
}

func UserPrint(u User) {
	fmt.Println(u.Age, u.LastName, u.FirstName)
}

func AdminPrint(a Admin) {
	fmt.Println(a.Email, a.Password, a.Age, a.FirstName, a.LastName)
}

func ClearName(u *User) {
	u.FirstName = ""
}

func NewUser(firstName string, lastname string, age int) User {
	return User{
		FirstName: firstName,
		LastName:  lastname,
		Age:       age,
	}
}

func NewAdmin(u User, email string, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
		},
	}
}

func ValidateUser(u User) error {
	if u.FirstName == "" || u.LastName == "" || u.Age < 20 {
		return errors.New("first and last Name can not be empty and age should be more than 20")
	}
	return nil
}
