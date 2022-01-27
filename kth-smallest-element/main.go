package main

import (
	"fmt"
	"sort"
)
//kth from the start == (len(arr) -k -1)th from the end 
func main() {
	var k int
	fmt.Println(" enter the value of k")
	fmt.Scanln(&k)
	arr := [...]int{12, 331, 31, 451, 54, 52} // unsorted
	arrSlice := arr[:]                        // created a slice of the array
	sort.Ints(arrSlice)
	fmt.Println(" the array after sorting becomes", arrSlice)
	fmt.Println("the ", k, "th smallest integer is ", arrSlice[k-1])

}

//  sort.Sort(sort.Reverse(sort.StringSlice(arr)))
