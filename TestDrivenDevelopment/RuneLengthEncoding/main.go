package main

import (
	"strconv"
	"strings"
)

var digit []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var alphabet []string = makeAlphabet()

func isDigit(input string) bool {
	for _, val := range digit {
		if val == input {
			return true
		}
	}

	return false
}

func isLetter(input string) bool {
	for _, val := range alphabet {
		if val == input {
			return true
		}
	}

	return false
}

func isSelectedSpecialChar(input string) bool {
	if input == "@" || input == "#" {
		return true
	}

	return false
}

func isWhiteSpace(input string) bool {
	return input == " "
}

func isAllowed(input string) bool {
	return !isDigit(input) && (isLetter(input) || isSelectedSpecialChar(input) || isWhiteSpace(input))
}

func Encode(input string) (encoded_msg string) {
	rep := 1
	upperCaseInput := strings.ToUpper(input)
	for index, val := range upperCaseInput {
		if isDigit(string(val)) {
			return "error: digits not allowed"
		} else if !isLetter(string(val)) && string(val) != " " && !isSelectedSpecialChar(string(val)) {
			return "error special character"
		} else {

			if index+1 < len(upperCaseInput) {
				if upperCaseInput[index] == upperCaseInput[index+1] {
					rep++
				} else {
					charCount := ""
					if rep > 1 {
						charCount = strconv.Itoa(rep)
					}
					encoded_msg = encoded_msg + charCount + string(upperCaseInput[index])
					rep = 1
				}
			} else {
				charCount := ""
				if rep > 1 {
					charCount = strconv.Itoa(rep)
				}
				encoded_msg = encoded_msg + charCount + string(upperCaseInput[index])
				rep = 1
			}

		}
	}
	//fmt.Println("encoded_msg: ", encoded_msg)
	return
}

func makeAlphabet() []string {
	var alphabet = make([]string, 0, 26)
	var ch byte
	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, string(ch))
	}

	return alphabet
}

//[0,3=] i+1=4
