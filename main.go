package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%v is even", number)
			println()
		} else {
			fmt.Printf("%v is odd", number)
			println()
		}
	}
}