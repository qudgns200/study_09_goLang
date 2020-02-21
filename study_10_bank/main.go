package main

import (
	"fmt"
	"study_10_bank/accounts"
)

func main() {
	account := accounts.NewAccount("paul")
	// account.Deposit(1000)
	// err := account.Withdraw(2000)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}