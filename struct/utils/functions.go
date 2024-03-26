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

func UserPrint(u User) {
	fmt.Println(u.Age, u.LastName, u.FirstName)
}

func ClearName(u *User) {
	u.FirstName = ""
}

func ValidateUser(u User) error {
	if u.FirstName == "" || u.LastName == "" || u.Age < 20 {
		return errors.New("first and last Name can not be empty and age should be more than 20")
	}
	return nil
}
