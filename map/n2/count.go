package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	wordsMap := make(map[string]int)

	punctuation := []string{".", ",", "!", "?", "-", "(", ")", "\""}
	for _, p := range punctuation {
		text = strings.ReplaceAll(text, p, "")
	}
	text = strings.ToLower(text)

	words := strings.Fields(text)

	for _, word := range words {
		wordsMap[word]++
	}

	return wordsMap
}

func main() {
	text := "Манул - это скрытный и редкий кот. Манул живёт в горах и степях. У манула густая шерсть. Манул красивый."

	wordCount := countWords(text)

	fmt.Println("Частота слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
