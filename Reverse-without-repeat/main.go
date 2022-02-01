package main

import "fmt"

/*
Problem #3

Given a string, say S, print it in reverse manner eliminating the repeated characters and spaces.

Example 1:

Input: S = "GEEKS FOR GEEKS" // SKEEG ROF SKEEG
Output: "SKEGROF"//"SKEGROF"

Example 2:

Input: S = "I AM AWESOME"
Output: "EMOSWAI"


*/

/*
solution


1.  reverse it
2.  make a slice/array without repeating characters

*/
//rune
func reverse(s string) (reverse string) {
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func IsContain(sample, char string) bool {
	if sample != "" {
		for _, value := range sample {
			if string(value) == char {
				return true
			}
		}
	}
	return false
}

func NoReapeat(s string) (result string) {
	rvrs := reverse(s)

	for _, value := range rvrs {
		if IsContain(result, string(value)) || string(value) == " " {
			continue
		} else {
			result += string(value)
		}
	}
	return
}

func main() {
	fmt.Println(NoReapeat("I AM AWESOME"))
}
