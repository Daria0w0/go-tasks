package main

import "fmt"

type Transport interface {
	Move()
	Stop()
}

type Car struct {
	Brand string
	Speed int
}

func (car Car) Move() {
	fmt.Printf("Автомобиль %s начинает движение со скоростью %d км/ч\n", car.Brand, car.Speed)
}

func (car Car) Stop() {
	fmt.Printf("Автомобиль %s остановился\n", car.Brand)
}

type Bicycle struct {
	Type string
}

func (bicycle Bicycle) Move() {
	fmt.Printf("%s велосипед начинает движение\n", bicycle.Type)
}

func (bicycle Bicycle) Stop() {
	fmt.Printf("%s велосипед остановился\n", bicycle.Type)
}

type Plane struct {
	Company string
	Height  int
}

func (plane Plane) Move() {
	fmt.Printf("Самолет авиакомпании %s поднялся на высоту %d\n", plane.Company, plane.Height)
}

func (plane Plane) Stop() {
	fmt.Printf("Самолет авиакомпании %s совершил посадку\n", plane.Company)
}

func printInfo(t Transport) {
	t.Move()
	t.Stop()
	fmt.Println("-----------------")
}

func main() {
	car := Car{"Toyota Camry", 60}
	bicycle := Bicycle{"Горный"}
	plane := Plane{"Аэрофлот", 10000}

	printInfo(car)
	printInfo(bicycle)
	printInfo(plane)
}
