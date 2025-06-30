package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("Сложение: %.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("Вычитание: %.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("Умножение: %.2f * %.2f = %.2f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("Деление: %.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Println("Деление на ноль невозможно!")
	}

	fmt.Printf("Остаток от деления: %.2f %% %.2f = %.2f\n", a, b, float64(int(a)%int(b)))
}
