// Go program that uses a function
// with variable argument list
package main

// Importing required packages
import (
	"fmt"
)

// Variadic function to return
// the sum of the numbers
func add(num ...int) int {
	sum := 0
	for j := range num {
		sum += j
	}
	return sum
}

func main() {
	fmt.Println("Sum =", add(1, 2, 3, 4, 5, 7, 8, 6, 5, 4))
}
