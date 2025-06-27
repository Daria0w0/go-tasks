package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Манул - это самый пушистый кот"
	arr := strings.Split(str, " ")

	fmt.Println(strings.Join(arr, "\n"))
}
