package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operation string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Ошибка: введите корректное числовое значение.")
		return
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scanln(&operation)

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Ошибка: введите корректное числовое значение.")
		return
	}

	result := 0.0
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}

	fmt.Printf("Результат: %v %s %v = %v\n", num1, operation, num2, result)
}
