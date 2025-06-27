package main

import "fmt"

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}

func main() {
	var year int

	fmt.Println("Год:")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Printf("Год %d - високосный", year)
	} else {
		fmt.Printf("Год %d - не високосный", year)
	}
}
