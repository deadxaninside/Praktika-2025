package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Grades    []int
}

// Метод для вычисления возраста
func (s Student) Age() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

// Метод для получения статуса
func (s Student) Status() string {
	average := 0.0
	for _, grade := range s.Grades {
		average += float64(grade)
	}
	average /= float64(len(s.Grades))
	switch {
	case average >= 4.5:
		return "отличник"
	case average >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}
func main() {
	student := Student{
		Name:      "София Афимина",
		BirthYear: 2005,
		Grades:    []int{5, 4, 5, 3},
	}
	fmt.Printf("%s, возраст: %d лет\n", student.Name, student.Age())
	fmt.Printf("Статус: %s\n", student.Status())
}
