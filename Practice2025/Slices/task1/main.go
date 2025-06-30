package main

import "fmt"

func main() {
	// Создаем пустой срез строк
	var slice []string

	// Динамически добавляем элементы
	slice = append(slice, "Первый")
	slice = append(slice, "Второй")
	slice = append(slice, "Третий")

	// Выводим все элементы
	fmt.Println("Содержимое среза:")
	for i, value := range slice {
		fmt.Printf("Индекс %d: %s\n", i, value)
	}
}
