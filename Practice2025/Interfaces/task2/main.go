package main

import "fmt"

// Интерфейс транспорта
type Transport interface {
	Move()
	Stop()
}

// Автомобиль
type Car struct{}

func (c Car) Move() {
	fmt.Println("Автомобиль едет")
}
func (c Car) Stop() {
	fmt.Println("Автомобиль остановился")
}

// Велосипед
type Bicycle struct{}

func (b Bicycle) Move() {
	fmt.Println("Велосипед катится")
}
func (b Bicycle) Stop() {
	fmt.Println("Велосипед остановился")
}
func useTransport(t Transport) {
	t.Move()
	t.Stop()
}
func main() {
	car := Car{}
	bike := Bicycle{}

	useTransport(car)
	useTransport(bike)
}
