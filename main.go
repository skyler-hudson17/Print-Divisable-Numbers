package main

import (
	"fmt"
)

func main() {
	var min, max, divisor int

	fmt.Print("Please enter a range: ")
	fmt.Scan(&min, &max)
	fmt.Print("Enter a Divisor: ")
	fmt.Scan(&divisor)

	for i := min; i <= max; i++ {
		if i%divisor == 0 {
			fmt.Print(i, ", ")
		}
	}
}
