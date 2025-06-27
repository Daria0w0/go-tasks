package main

import "fmt"

type BankAcc struct {
	Name    string
	Balance float64
}

func (balance *BankAcc) topUp(amount float64) {
	if amount > 0 {
		balance.Balance += amount
		fmt.Printf("Пополнение на %.2f руб\n", amount)
	} else {
		fmt.Println("Ошибка: сумма пополнения должна быть положительной")
	}
}

func (balance *BankAcc) withdraw(amount float64) {
	if amount > 0 {
		if balance.Balance >= amount {
			balance.Balance -= amount
			fmt.Printf("Снятие %.2f руб\n", amount)
		} else {
			fmt.Println("Ошибка: не достаточно средств на счете")
		}
	} else {
		fmt.Println("Ошибка: сумма снятия должна быть положительной")
	}
}

func (balance *BankAcc) getBalance() float64 {
	return balance.Balance
}

func main() {
	myAcc := BankAcc{"Sitnikova Daria", 7050}
	fmt.Println("Текущий баланс:", myAcc.getBalance())
	myAcc.topUp(19900)
	fmt.Println("Текущий баланс:", myAcc.getBalance())
	myAcc.withdraw(3100)
	fmt.Println("Текущий баланс:", myAcc.getBalance())
	myAcc.withdraw(30000)
	fmt.Println("Текущий баланс:", myAcc.getBalance())
}
