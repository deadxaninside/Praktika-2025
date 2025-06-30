package main

import "fmt"

func main() {
	// Целочисленные типы
	var a int = -42
	var b int8 = -127
	var c int16 = 32767
	var d int32 = -2147483648
	var e int64 = 9223372036854775807

	var f uint = 42
	var g uint8 = 255
	var h uint16 = 65535
	var i uint32 = 4294967295
	var j uint64 = 18446744073709551615

	// Числа с плавающей точкой
	var k float32 = 3.14
	var l float64 = 2.71828

	// Комплексные числа
	var m complex64 = 1 + 2i
	var n complex128 = 3 + 4i

	// Логический тип
	var o bool = true

	// Строки
	var p string = "Привет"

	// Байт
	var q byte = 65

	// Руна
	var r rune = 'Я'

	fmt.Println("Целые (со знаком):", a, b, c, d, e)
	fmt.Println("Целые (без знака):", f, g, h, i, j)
	fmt.Println("Дробные:", k, l)
	fmt.Println("Комплексные:", m, n)
	fmt.Println("Логическое:", o)
	fmt.Println("Строка:", p)
	fmt.Println("Байт (как символ):", string(q))
	fmt.Println("Руна (как символ):", string(r))
}
