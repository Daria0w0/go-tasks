package main

import "fmt"

func main() {
	var slice []string

	slice = append(slice, "Манул")
	slice = append(slice, "-")
	slice = append(slice, "пушистый")
	slice = append(slice, "кот")

	fmt.Println(slice)
}
