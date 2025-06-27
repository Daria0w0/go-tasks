package main

import (
	"fmt"
	"sort"
)

func search(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

func sorting(slice []int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	sort.Ints(copySlice)
	return copySlice
}

func filtering(slice []int, condition func(int) bool) []int {
	filtered := make([]int, 0)
	for _, v := range slice {
		if condition(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	num := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	if index, found := search(num, 7); found {
		fmt.Printf("Элемент 7 найден на позиции %d\n", index)
	} else {
		fmt.Println("Элемент 7 не найден")
	}

	fmt.Printf("Отсортированный по возрастанию срез: %v\n", sorting(num))

	filter := filtering(num, func(n int) bool {
		return n > 5
	})
	fmt.Printf("Числа больше 5: %v\n", filter)
}
