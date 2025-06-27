package main

import (
	"fmt"
)

func findLongestString(str ...string) string {
	if len(str) == 0 {
		return "Список пуст"
	}
	longest := str[0]
	for _, s := range str {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	fmt.Println(findLongestString("кот", "манул", "пушистый"))
	fmt.Println(findLongestString())
}
