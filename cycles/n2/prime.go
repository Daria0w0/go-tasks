package main

import "fmt"

func main() {
	for n := 2; n <= 100; n++ {
		isPrime := true

		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
			}
		}

		if isPrime {
			fmt.Printf("%d, ", n)
		}
	}
}
