package main

import "fmt"

// Функция с передачей по значению (копия)
func byValue(val int) {
	val = 100
	fmt.Println("Inside byValue:", val)
}

// Функция с передачей по указателю (оригинал)
func byPointer(ptr *int) {
	*ptr = 100
	fmt.Println("Inside byPointer:", *ptr)
}

func main() {
	original := 50

	// Передача по значению
	byValue(original)
	fmt.Println("After byValue:", original) // Не изменилось

	// Передача по указателю
	byPointer(&original)
	fmt.Println("After byPointer:", original) // Изменилось
}
