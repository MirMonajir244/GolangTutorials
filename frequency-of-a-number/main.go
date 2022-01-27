package main

import "fmt"

func main() {
	var input int
	count := 0
	arr := []int{1, 3, 2, 42, 53, 4, 5, 6, 3, 4, 5, 6, 3, 2, 1, 4, 5, 6, 7, 8, 9, 0, 2, 3, 4, 5, 67, 8, 9, 7, 5, 31, 4, 6, 5, 4, 6, 2, 5, 6, 5, 9, 56, 7, 4, 5, 6, 8, 3, 2, 4, 5, 8, 9, 1, 0}

	fmt.Println("Which value to count?")
	fmt.Scanln(&input)

	for _, value := range arr {
		if value == input {
			count++
		}
	}

	fmt.Println("the occurence of ", input, "in the array is ", count, "times")

}
