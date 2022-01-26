package main

import "fmt"

//var class map[int]string

func main() {
	student_class := map[int]string{
		1: "Aaditya",
		2: "Sharanya",
		3: "Devanshu",
	}
	student_class2 := map[int]string{
		21: "Gitanshu",
		22: "Saraansh",
		35: "Tarun",
	}
	student_class[15] = "Vishnu"
	student_class[17] = "Praveen"

	for key, value := range student_class2 {
		student_class[key] = value
		delete(student_class2, key)
		//student_class[21]="Gitanshu"
	}

	var new_map = make(map[int]string)

	for key, value := range student_class {
		new_map[key] = value
	}
	new_map[15] = "Rohan"

	fmt.Println("the original map")
	for roll_no, name := range student_class {
		fmt.Println("Roll no: ", roll_no, " Name: ", name)
	}

	fmt.Println("the new map")
	for roll_no, name := range new_map {
		fmt.Println("Roll no: ", roll_no, " Name: ", name)
	}
	/*	student, ok := student_class[15]
		if ok {
			fmt.Println(student)
		} else {
			fmt.Println("Key invalid")
		}
	*/
	/*
		delete(student_class, 15)

		for roll_no, name := range student_class {
			fmt.Println("Roll no: ", roll_no, " Name: ", name)
		}
	*/
	//sort.Ints()

}

/*
arr[0] ;2020
arr[1]; 2021
arr[2]; 2022
*/
/*
arr[100]

100th = 1->2->3>......>100

key:Value
"RollNo":101
*/

/*
map[Key_Type]Value_Type{}


var mymap map [int] string
var mymap map[int]string
pet_name, ok := m_a_p[90]
delete(m_a_p, 90)
new_map := m_a_p

    // Perform modification in new_map
    new_map[96] = "Parrot"
    new_map[98] = "Pig"
*/
/*
R->row (input by user)
C-> col (input by user)
arr1 := make([][]int, R)

for i := 0; i < R; i++{
	  arr1[i] = make([]int, C)
 }

 arr1={{0,0,0,0},{0,0,0},{0,0,0,},{0,0,0}}
*/
