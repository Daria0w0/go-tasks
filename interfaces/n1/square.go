package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Square() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (s Rectangle) Square() float64 {
	return s.Height * s.Width
}

type Circle struct {
	Radius float64
}

func (s Circle) Square() float64 {
	return math.Pi * s.Radius * s.Radius
}

func printSquare(shape Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", shape.Square())
}

func main() {
	rectangle := Rectangle{5, 3}
	circle := Circle{4}

	fmt.Println("Прямоугольник:")
	printSquare(rectangle)

	fmt.Println("\nКруг:")
	printSquare(circle)
}
