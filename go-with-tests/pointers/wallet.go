package pointers

import (
	"errors"
	"fmt"
)

type Cash int

type Account interface {
	Balance() int
	Deposit(int)
}

type Wallet struct {
	balance Cash
}

func (w *Wallet) Balance() Cash {
	return	w.balance
}

func (w *Wallet) Deposit(amount Cash) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Cash) error {
	if w.balance < amount {
		return errors.New("insufficient balance")
	}
	w.balance -= amount
	return nil
}

func (c Cash) String() string {
	return fmt.Sprintf("PHP %d", c)
}