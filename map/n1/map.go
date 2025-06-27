package main

import "fmt"

func add(Map map[string]int, name string, mark int) {
	Map[name] = mark
}

func search(Map map[string]int, name string) int {
	return Map[name]
}

func remove(Map map[string]int, name string) {
	delete(Map, name)
}

func main() {
	markMap := make(map[string]int)
	fmt.Println("Исходная карта:", markMap)

	add(markMap, "Daria", 5)
	add(markMap, "Anton", 4)
	fmt.Println("Функция добавления:", markMap)

	fmt.Println("Функция поиска по ключу Daria:", search(markMap, "Daria"))

	remove(markMap, "Daria")
	fmt.Println("Функция удаления по ключу Daria:", markMap)
}
