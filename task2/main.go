package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	hD := 0

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hD++
			}
		}

		return hD, nil
	}

	return hD, errors.New("both strings are  not equal")
	panic("Implement the Distance function")
}

func main() {
	//var a, b string
	// My roll number is 21. My age is 23.
	fmt.Println(" My roll no is ", 21, "my age is ", 23)
	fmt.Println(fmt.Sprintf("My roll no is %d. My age is %d", 21, 23))
	//fmt.Scanln(&a)

	fmt.Println(" input DNA strand no. 2")
	//fmt.Scanln(&b)
	//dist, err := Distance(a, b)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println("the hamming distance is ", dist)
	//	}

}
