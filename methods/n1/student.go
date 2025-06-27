package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Course    int
	GPA       float64
}

func (student Student) getAge() int {
	currentYear := time.Now().Year()
	return currentYear - student.BirthYear
}

func (student Student) getStatus() string {
	switch {
	case student.GPA >= 4.5:
		return "отличник"
	case student.GPA >= 4.0:
		return "хорошист"
	default:
		return "троечник"
	}
}

func main() {
	student := Student{"Daria", 2005, 2, 4.5}
	fmt.Printf("%s:\n", student.Name)
	fmt.Printf("Возраст: %d лет\n", student.getAge())
	fmt.Printf("Статус: %s\n", student.getStatus())
}
