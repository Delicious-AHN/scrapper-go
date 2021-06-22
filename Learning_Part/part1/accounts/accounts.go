package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errnoMoney = errors.New("Can't Withdraw Fucking poor man")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (theAccount *Account) Deposit(amount int) {
	theAccount.balance += amount
}

func (theAccount Account) Balance() int {
	return theAccount.balance
}

func (theAccount *Account) Withdraw(amount int) error {
	if theAccount.balance < amount {
		return errnoMoney
	}
	theAccount.balance -= amount
	return nil
}
