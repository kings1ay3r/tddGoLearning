package pointers

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("Insufficient Balance")

type USD int
type Wallet struct {
	balance USD
}

type stringer interface {
	BalaceString() string
}

func (d Wallet) BalanceString() string {
	return "USD " + fmt.Sprint(d.balance) + ",000"
}
func (w *Wallet) Deposit(amount USD) {
	w.balance += amount

}
func (w *Wallet) Withdraw(amount USD) error {
	if w.balance > amount {
		w.balance -= amount
	} else {
		return errors.New("Insufficient Balance")
	}
	return nil

}
func (w *Wallet) Balance() string {
	return w.BalanceString()
}
