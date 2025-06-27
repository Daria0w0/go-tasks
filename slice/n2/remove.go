package main

import "fmt"

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	mySlice := []string{"Манул", "-", "пушистый", "кот"}

	fmt.Println(remove(mySlice, 1))
}
