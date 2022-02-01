package main

/*
Banking Application

1.  log in
*/

/*
1. let the user enter username,password
2. verify it
3. if username and pass is not provided, use fingerprint
*/
func Factorial(input int) int {
	fact := 1
	if input == 0 {
		return 1
	} else if input == 1 {
		return 1
	} else if input < 0 {
		return 0
	}

	for i := input; i > 0; i-- {
		fact *= i
	}

	return fact

}

func main() {

}
