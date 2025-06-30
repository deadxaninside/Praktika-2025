package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	words := strings.Fields(text)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[strings.ToLower(word)]++
	}

	return frequency
}

func main() {
	text := "Go это язык Go который язык программирования"
	freq := wordFrequency(text)

	fmt.Println("Частота слов:")
	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
