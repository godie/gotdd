package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return (*wallet).balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficientFunds
	}
	wallet.balance -= amount
	return nil

}

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}
