package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

// Функция для создания нового студента
func NewStudent(name string, age, course int, avgGrade float64) Student {
	return Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

// Функция для вывода информации о студенте
func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n",
		s.Name, s.Age, s.Course, s.AvgGrade)
}

// Функция для повышения курса
func (s *Student) Promote() {
	s.Course++
}
func main() {
	// Создаем студента
	student := NewStudent("София Афимина", 20, 2, 4.3)
	student.PrintInfo()
	// Повышаем курс
	student.Promote()
	fmt.Println("\nПосле повышения:")
	student.PrintInfo()
}
