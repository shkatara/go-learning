package main

import (
	"fmt"
)

func bankMenu() {
	fmt.Println("\nWelcome to the bank")
	fmt.Println("1) Check Balance")
	fmt.Println("2) Add Balance")
	fmt.Println("3) Withdraw Balance")
	fmt.Println("4) Exit\n")
	fmt.Print("What do you need: ")
}

func addBalance(balance float32, to_add float32) float32 {
	fmt.Print("\nHow much to add to Balance: ")
	fmt.Scan(&to_add)
	updated_balance := balance + to_add
	return updated_balance
}

func removeBalance(balance float32, to_del float32) float32 {
	fmt.Print("\nHow much to Withdraw: \n")
	fmt.Scan(&to_del)
	if to_del > balance {
		fmt.Println("Not Enough Funds. Can not Withdraw.")
		return 0
	} else {
		updated_balance := balance - to_del
		return updated_balance
	}
}

func bank_start(balance float32) {
	var choice int
	var to_add float32
	var to_del float32

	for true {
		bankMenu()
		fmt.Scan(&choice)
		wantsCheckBalance := choice == 1
		wantsAddBalance := choice == 2
		wantsDelBalance := choice == 3
		wantsExit := choice == 4
		if wantsCheckBalance {
			fmt.Println("\nBalance is", balance)
		} else if wantsAddBalance {
			balance = addBalance(balance, to_add)
			fmt.Println("\nAdd Successful. Updated Balance is", balance)
		} else if wantsDelBalance {
			reduced_balance := removeBalance(balance, to_del)
			if reduced_balance != 0 {
				fmt.Println("\nWithdrawal Successful. Updated Balance is", reduced_balance)
				balance = reduced_balance
			}
		} else if wantsExit {
			fmt.Println("\nExiting Bank")
			break
		}
	}
}

func main() {
	var balance float32 = 1000

	fmt.Println("Initializing bank...")
	bank_start(balance)
}
