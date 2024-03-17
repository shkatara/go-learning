package main

import (
	"fmt"
	"os"

	"example.com/bank/utils"
)

var choice int
var to_add float32
var balanceFile = "balance.txt"
var to_del float32

func bank_start(balance float32) {
bankWorkLoop:
	for {
		utils.BankMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\nBalance is", balance)
		case 2:
			balance = utils.AddBalance(balance, to_add)
			fmt.Println("\nAdd Successful. Updated Balance is", balance)
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
	data, err := os.ReadFile(balanceFile)
	utils.Check(err)
	balance := string(data)
	fmt.Print(balance)
}
