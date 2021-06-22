package main

import (
	"fmt"
	"log"

	"github.io/Delicious-Ahn/nomad-go/Learning_Part/part1/accounts"
)

func main() {
	account := accounts.NewAccount("Hyunjae")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
