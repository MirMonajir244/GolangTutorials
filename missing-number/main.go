package main

import "fmt"

func findSmallestMissingPositive(arr []int) int {

	for i := 1; i < len(arr); i++ {
		found := false
		fmt.Println("checking ", i)
		for _, value := range arr {
			if value == i {
				found = true
				break
			}
		}

		if !found {
			return i
		}

	}

	return arr[len(arr)-1]

}

func main() {
	arr := []int{1, 1, 0, -1, -2}

	fmt.Println("the smallest missing number is ", findSmallestMissingPositive(arr))
}
