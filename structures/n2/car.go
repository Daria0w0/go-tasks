package main

import "fmt"

type Engine struct {
	Type      string  //Тип двигателя
	Volume    float64 //Объем
	Power     float64 //Мощность
	IsRunning bool    //Запущен ли двигатель
}

func printInfoEngine(engine *Engine) {
	fmt.Printf("Тип двигателя: %s\n", engine.Type)
	fmt.Printf("Объем двигателя: %f\n", engine.Volume)
	fmt.Printf("Мощность двигателя: %f\n", engine.Power)
	status := "не запущен"
	if engine.IsRunning {
		status = "запущен"
	}
	fmt.Printf("Двигатель: %s\n", status)
}

type Car struct {
	Brand   string //Марка
	Year    int    //Год выпуска
	Mileage int    //Пробег
	engine  Engine //Двигатель
}

func printInfoCar(car *Car) {
	fmt.Printf("Марка: %s\n", car.Brand)
	fmt.Printf("Год: %d\n", car.Year)
	fmt.Printf("Пробег: %d км\n", car.Mileage)
	fmt.Println("\nДвигатель:")
	printInfoEngine(&car.engine)
}

func start(engine *Engine) {
	if !engine.IsRunning {
		engine.IsRunning = true
		fmt.Println("\nДвигатель запущен")
	} else {
		fmt.Println("\nДвигатель уже работает")
	}
}

func stop(engine *Engine) {
	if engine.IsRunning {
		engine.IsRunning = false
		fmt.Println("\nДвигатель выключен")
	} else {
		fmt.Println("\nДвигатель уже выключен")
	}
}

func main() {
	engine := Engine{"бензиновый", 150, 2.0, false}
	car := Car{"Toyota", 2020, 35000, engine}

	printInfoCar(&car)
	start(&car.engine)
	printInfoEngine(&car.engine)
	start(&car.engine)
	stop(&car.engine)
}
