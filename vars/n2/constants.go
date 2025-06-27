package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	const e = math.E

	r := 2.0
	C := 2 * pi * r

	logE := math.Log(e)

	fmt.Println("Число pi =", pi)
	fmt.Println("Длина окружности с радиусом 2 =", C)
	fmt.Println("Число e =", e)
	fmt.Println("Натуральные логарифм от e =", logE)
}
