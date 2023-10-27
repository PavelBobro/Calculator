package main

import (
	"fmt"
)

func main() {
	var a, c float64
	var b string

	fmt.Println("Введите первое число")
	fmt.Scan(&a)
	fmt.Println("Укажите что сделать с числом")
	fmt.Scan(&b)
	fmt.Println("Введите второе число")
	fmt.Scan(&c)

	switch b {
	case "+":
		result := a + c
		fmt.Printf("%.2f + %.2f = %.2f\n", a, c, result)
	case "-":
		result := a - c
		fmt.Printf("%.2f + %.2f = %.2f\n\n", a, c, result)
	case "*":
		result := a * c
		fmt.Printf("%.2f + %.2f = %.2f\n", a, c, result)

	case "/":
		if c == 0 {
			fmt.Println("Деление на ноль недопустимо")
		} else {
			result := a / c
			fmt.Printf("%.2f + %.2f = %.2f\n", a, c, result)
		}
	default:
		fmt.Println("ввели не правельный символ, требуется указать только + - * или / ")
	}

}
