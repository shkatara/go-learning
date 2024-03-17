package utils

import (
	"fmt"
	"os"
	"strconv"
)

func BankMenu() {
	fmt.Println("Initializing bank...")
	fmt.Println("\nWelcome to the bank")
	fmt.Println("1) Check Balance")
	fmt.Println("2) Add Balance")
	fmt.Println("3) Withdraw Balance")
	fmt.Println("4) Exit\n")
	fmt.Print("What do you need: ")
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteBalanceToFile(balance_filename string, balance float64) {
	balanceText := fmt.Sprint(balance) // convert float to string so we can convert string to []byte in os.Writefile func
	os.WriteFile(balance_filename, []byte(balanceText), 0644)
}

func ReadBalanceFromFile(balance_filename string) float64 {
	data, err := os.ReadFile(balance_filename)
	Check(err)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func AddBalance(balance float64, to_add float64) float64 {
	fmt.Print("\nHow much to add to Balance: ")
	fmt.Scan(&to_add)
	updated_balance := balance + to_add
	return updated_balance
}

func RemoveBalance(balance float64, to_del float64) float64 {
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
