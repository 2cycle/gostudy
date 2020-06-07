package accounts

import (
	"errors"
	"fmt"
)

// Account struct
// lowerClass start var is private if use publie use Upperclass staring word
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// now balance of yours
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount of you account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//Change Owner
func (a *Account) ChangeOwner(newOwner string) {
	defer fmt.Println("the owner changed to ", newOwner)
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
