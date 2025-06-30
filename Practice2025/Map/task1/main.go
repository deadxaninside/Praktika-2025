package main

import "fmt"

func main() {
	// Создаем карту для хранения оценок
	grades := make(map[string]int)

	// Добавление оценок
	grades["Алексей"] = 5
	grades["Мария"] = 4
	grades["Иван"] = 3

	// Поиск оценки
	name := "Мария"
	if grade, exists := grades[name]; exists {
		fmt.Printf("%s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}

	// Удаление оценки
	delete(grades, "Иван")

	// Вывод всех оценок
	fmt.Println("Текущие оценки:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
}
