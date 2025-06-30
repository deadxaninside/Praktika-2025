package main

import (
	"fmt"
	"math"
)

// Интерфейс формы
type Shape interface {
	Area() float64
}

// Прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Круг
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Площадь: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	printArea(rect)
	printArea(circle)
}
