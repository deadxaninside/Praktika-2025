package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b float64
	var operator string
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
			os.Exit(1)
		}
		result = a / b
	default:
		fmt.Println("Неверный оператор!")
		os.Exit(1)
	}
	fmt.Printf("Результат: %.2f\n", result)
}
