package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Course int
	GPA    float64
}

func newStudent(name string, age int, course int, gpa float64) Student {
	newStudent := Student{name, age, course, gpa}
	return newStudent
}

func printInfo(student Student) {
	fmt.Printf("Студент: %s\n", student.Name)
	fmt.Printf("Возраст: %d\n", student.Age)
	fmt.Printf("Курс: %d\n", student.Course)
	fmt.Printf("Средний балл: %f\n\n", student.GPA)
}

func courseUp(student Student) {
	student.Course++
	fmt.Printf("Курс студента %s повышен до %d\n\n", student.Name, student.Course)
}

func main() {
	student1 := newStudent("Daria", 20, 2, 4.5)
	student2 := newStudent("Motya", 18, 1, 4.9)
	printInfo(student1)
	printInfo(student2)
	courseUp(student2)
}
