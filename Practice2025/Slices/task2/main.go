package main

import "fmt"

func removeElement(slice []string, index int) []string {
	// Проверка корректности индекса
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Удаляем элемент, объединяя части до и после индекса
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Исходный срез
	fruits := []string{"Яблоко", "Банан", "Апельсин", "Груша"}

	fmt.Println("До удаления:", fruits)

	// Удаляем элемент с индексом 1 (Банан)
	fruits = removeElement(fruits, 1)

	fmt.Println("После удаления:", fruits)
}
