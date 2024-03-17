package main

import (
	"fmt"

	"example.com/bank/utils"
)

var choice int
var to_add float64

const balanceFile string = "balance.txt"

var to_del float64

func bank_start(balance float64) {
bankWorkLoop:
	for {
		balance_from_file := utils.ReadBalanceFromFile(balanceFile)

		utils.BankMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("\nBalance is", balance_from_file)
		case 2:
			updated_balance := utils.AddBalance(balance, to_add)
			fmt.Println("\nAdd Successful. Updated Balance is", updated_balance)
			utils.WriteBalanceToFile(balanceFile, updated_balance)
		case 3:
			reduced_balance := utils.RemoveBalance(balance, to_del)
			if reduced_balance != 0 {
				fmt.Println("\nWithdrawal Successful. Updated Balance is", reduced_balance)
				balance = reduced_balance
			}
		case 4:
			fmt.Println("\nExiting Bank")
			fmt.Println("Thanks for choosing our Bank.")
			break bankWorkLoop
		default:
			fmt.Println("\nWrong Input. Exiting")
			fmt.Println("Thanks for choosing our Bank.")
			break bankWorkLoop
		}
	}
}

func main() {
	balance := utils.ReadBalanceFromFile(balanceFile)
	bank_start(balance)
}
