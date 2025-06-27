package main

import "fmt"

func main() {
	var num1, num2 float32
	var op string

	fmt.Println("Первое число:")
	fmt.Scan(&num1)
	fmt.Println("Операция (+, -, *, %):")
	fmt.Scan(&op)
	fmt.Println("Второе число:")
	fmt.Scan(&num2)

	switch op {
	case "+":
		fmt.Println("Результат:", num1+num2)
	case "-":
		fmt.Println("Результат:", num1-num2)
	case "*":
		fmt.Println("Результат:", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка. Деление на 0")
		} else {
			fmt.Println("Результат:", num1/num2)
		}
	default:
		fmt.Println("Ошибка. Неизвестный оператор")
	}
}
