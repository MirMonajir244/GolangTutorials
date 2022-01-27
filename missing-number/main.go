package main

import "fmt"

/*
You are given an array arr[] of N integers including 0. The task is to find the smallest positive number missing from the array.

Example 1:

Input:
N = 5
arr[] = {1,3,5,2,4}
Output: 6
Explanation: Smallest positive missing
number is 6.
Example 2:

Input:
N = 5
arr[] = {0,-10,1,-20,3}
Output: 2
Explanation: Smallest positive missing
number is 2.

*/
func findSmallestMissingPositive(arr []int) int {

	for i := 1; i < len(arr); i++ {
		found := false
		//fmt.Println("checking ", i)
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

	return len(arr) + 1

}

func main() {
	arr := []int{1, 3, 5, 2, 4}

	fmt.Println("the smallest missing number is ", findSmallestMissingPositive(arr))
}
