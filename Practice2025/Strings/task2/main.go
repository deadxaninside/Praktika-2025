package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Go это современный язык программирования"

	// Разбиваем строку на слова
	words := strings.Fields(sentence)

	// Выводим каждое слово на новой строке
	for i, word := range words {
		fmt.Printf("Слово %d: %s\n", i+1, word)
	}
}
