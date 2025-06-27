package main

import "fmt"

func add(x int) {
	x = x + 1
	fmt.Println("Внутри функции с предачей по значению:", x)
}

func addPoint(x *int) {
	*x = *x + 1
	fmt.Println("Внутри функции с предачей по указателю:", *x)
}

func main() {
	num := 100
	fmt.Println("До вызова add():", num)
	add(num)
	fmt.Printf("После вызова add():%d\n\n", num)

	fmt.Println("До вызова addPoint():", num)
	addPoint(&num)
	fmt.Printf("После вызова addPoint():%d\n", num)
}
