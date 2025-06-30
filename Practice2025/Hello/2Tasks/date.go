package main

import (
	"fmt"
	"time"
)

func main() {
	name := "София"
	currentDate := time.Now().Format("02.01.2006")

	message := fmt.Sprintf("Привет, меня зовут %s. Сегодня %s", name, currentDate)
	fmt.Println(message)
}
