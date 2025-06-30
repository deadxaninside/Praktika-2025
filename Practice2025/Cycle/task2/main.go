package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Простые числа до 100:")
	for num := 2; num <= 100; num++ {
		if isPrime(num) {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}
