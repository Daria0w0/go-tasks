package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	fmt.Printf("До обмена:\nx = %d, y = %d\n", x, y)

	swap(&x, &y)

	fmt.Printf("После обмена:\nx = %d, y = %d\n", x, y)
}
