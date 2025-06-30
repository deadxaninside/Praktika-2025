package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

// Метод для пополнения счёта
func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	acc.balance += amount
	return nil
}

// Метод для снятия средств
func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств")
	}
	acc.balance -= amount
	return nil
}

// Метод для проверки баланса
func (acc BankAccount) Balance() float64 {
	return acc.balance
}

func main() {
	account := BankAccount{owner: "Петр Петров", balance: 1000}

	// Пополнение
	if err := account.Deposit(500); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Счёт пополнен")
	}

	// Снятие
	if err := account.Withdraw(200); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Средства сняты")
	}

	// Проверка баланса
	fmt.Printf("Текущий баланс: %.2f руб.\n", account.Balance())
}
