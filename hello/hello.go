package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Daria"
	fmt.Println("Hello,", name)
	fmt.Println("Today is", time.Now().Format("02.01.2006"))
}
