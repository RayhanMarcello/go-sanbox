package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	OwnerName string
	Balance   int
}

// var wallet []*Wallet

func (w *Wallet) Deposito(ammount int) {
	total := w.Balance + ammount
	fmt.Printf("saldo anda: %v dan deposito anda : %v \n" , w.Balance, ammount)
	fmt.Printf("total saldo : %v \n", total)
	
}

func newWallet() *Wallet {
	return &Wallet{
		OwnerName: "rayhan",
		Balance:   2000000,
	}
}

func (w *Wallet) Withdraw(ammount int)error{
	cukupGa := w.Balance - ammount
	if w.Balance < ammount{
		return errors.New("saldo ga cukup bro")	
	}

		fmt.Printf("berhasil saldo anda : %v \n", cukupGa)
	return nil
}

