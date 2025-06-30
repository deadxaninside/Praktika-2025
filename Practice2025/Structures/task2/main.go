package main

import "fmt"

type Engine struct {
	Power    int
	Volume   float64
	FuelType string
}
type Car struct {
	Brand  string
	Model  string
	Year   int
	Engine Engine
}

// Функция для вывода информации об автомобиле
func (c Car) PrintInfo() {
	fmt.Printf("Автомобиль: %s %s (%d г.)\n", c.Brand, c.Model, c.Year)
	fmt.Printf("Двигатель: %.1f л, %d л.с., топливо: %s\n",
		c.Engine.Volume, c.Engine.Power, c.Engine.FuelType)
}
func main() {
	// Создаем двигатель
	engine := Engine{
		Power:    150,
		Volume:   1.8,
		FuelType: "бензин",
	}
	// Создаем автомобиль с этим двигателем
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Year:   2020,
		Engine: engine,
	}
	car.PrintInfo()
	// Модифицируем мощность двигателя
	car.Engine.Power = 180
	fmt.Println("\nПосле тюнинга:")
	car.PrintInfo()
}
