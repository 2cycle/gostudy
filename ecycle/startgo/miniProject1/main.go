package main

import (
	"fmt"
	"log"

	"github.com/ecycle/startgo/miniProject1/accounts"
)

func main() {
	/*
		account := banking.Account{Owner: "wally", Balance: 10000}
		fmt.Println(account)
	*/
	account := accounts.NewAccount("wally")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("newWally")
	fmt.Println(account) // call String()
}
