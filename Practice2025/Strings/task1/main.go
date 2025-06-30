package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World! GoLang"

	// Подсчет символов
	length := len(str)
	fmt.Printf("Длина строки: %d символов\n", length)

	// Поиск подстроки
	substr := "Go"
	contains := strings.Contains(str, substr)
	fmt.Printf("Содержит '%s': %t\n", substr, contains)

	// Изменение регистра
	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	fmt.Println("В нижнем регистре:", lower)
	fmt.Println("В верхнем регистре:", upper)
}
