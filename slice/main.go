package main

import "fmt"



func main() {
	/*
		arr[0] = 1
		arr[1] = 3
		arr[2] = 4
	*/
	// Creating an array whose size is determined
	// by the number of elements present in it
	// Using ellipsis
	//arr := [...]int{1, 2, 3, 4, 5, 4, 5, 7, 4, 2, 2}

	//arr1 := [2][2]int{{1, 2}, {3, 4}}

	/*myarray := [...]int{29, 79, 49, 39,
	20, 49, 48, 49}
	*/
	// Iterate array using for loop
	//for x := 0; x < len(myarray); x++ {
	//		fmt.Printf("%d\n", myarray[x])
	//	}

	/*	for index, value := range myarray {
			fmt.Println("element at ", index, ": ", value)
		}
	*/
	/*
		myarr := [2][2]int{{1, 2}, {3, 4}}

		for index1, value1 := range myarr {
			for index2, value2 := range value1 {
				fmt.Println("element at ", index1, " ", index2, " is", value2)
			}
		}
	*/

	// SLICE
	//var s []int
	//s := make([]int, 5)

	//make([]data_type, length, capacity)
	/*for index := range s {
		s[index] = 5
	}
	*/
	/*
		for i := 0; i < 500; i++ {
			s = append(s, 20)
		}

		s = append(s, 25)
		s = append(s, 23)
		s = append(s, 26)

		fmt.Println(s)

	*/
	//making a slice out of an array
	arr := []int{0, 3, 2, 5, 6, 4}
	arr1 := arr[0:5] // from arr[0] to arr[4] ( excluding everything arr[5] onwards)
	// arr1 = {0,3,2,5,6,4}
	arr1[0] = 10 // it will change the original slice as well
	// Strings are immutable
	//st := "Aastrungtskjlalashlk"

	/*	for _, value := range st {
			fmt.Println(string(value))
		}
	*/
	//strings.Contains(st,"a")
	//fmt.Println(strings.Count("test", "t"))
	/*  s.HasPrefix("test", "te"))
	    p("HasSuffix: ", s.HasSuffix("test", "st"))
	    p("Index:     ", s.Index("test", "e"))
	    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	    p("Repeat:    ", s.Repeat("a", 5))
	    p("Replace:   ", s.Replace("foo", "o", "0", -1))
	    p("Replace:   ", s.Replace("foo", "o", "0", 1))
	    p("Split:     ", s.Split("a-b-c-d-e", "-"))
	    p("ToLower:   ", s.ToLower("TEST"))
	    p("ToUpper:   ", s.ToUpper("test"))
	    p()
	*/

	//res1 := strings.TrimSpace("  tsssTarun  ")
	//res1 := strings.TrimSuffix("tsssTaruntheGolangGuy", "GolangGuy")

	//res1 := strings.TrimPrefix("tsssTaruntheGolangGuy", "tsss")

	//fmt.Println(res1)
	//fmt.Println(arr)

	// Stringer
	/* A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
	 */
	

	//fmt.Println("name: ", a.Name, "age: ", a.Age)
	}
