package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Первое число:")
	fmt.Scan(&num1)
	fmt.Println("Второе число:")
	fmt.Scan(&num2)

	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	if num2 != 0 {
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
		fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)
	} else {
		fmt.Println("Ошибка. Деление на 0")
	}
}
