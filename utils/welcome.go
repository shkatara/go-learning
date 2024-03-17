package utils

import "fmt"

func BankMenu() {
	fmt.Println("Initializing bank...")
	fmt.Println("\nWelcome to the bank")
	fmt.Println("1) Check Balance")
	fmt.Println("2) Add Balance")
	fmt.Println("3) Withdraw Balance")
	fmt.Println("4) Exit\n")
	fmt.Print("What do you need: ")
}

func AddBalance(balance float32, to_add float32) float32 {
	fmt.Print("\nHow much to add to Balance: ")
	fmt.Scan(&to_add)
	updated_balance := balance + to_add
	return updated_balance
}

func RemoveBalance(balance float32, to_del float32) float32 {
	fmt.Print("\nHow much to Withdraw: ")
	fmt.Scan(&to_del)
	if to_del > balance {
		fmt.Println("Not Enough Funds. Can not Withdraw.")
		return 0
	} else {
		updated_balance := balance - to_del
		return updated_balance
	}
}
