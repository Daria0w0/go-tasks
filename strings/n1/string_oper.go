package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "maNul0W0manul"
	var substr = "manul"

	fmt.Printf("1. Подсчет символов в строке %s: %d\n\n", str, len(str))
	fmt.Printf("2. Поиск подстроки %s в строке %s:\nВ начале: %t\nВ конце: %t\nВсего: %d\n\n", substr, str, strings.HasPrefix(str, substr), strings.HasSuffix(str, substr), strings.Count(str, substr))
	fmt.Printf("3. Изменение регистра: %s, %s", strings.ToUpper(str), strings.ToLower(str))
}
