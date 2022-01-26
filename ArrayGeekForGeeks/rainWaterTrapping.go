package main

import "fmt"

var arr = []int{3, 0, 0, 4, 0, 4}
var maxHieght int

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func volume(h int) int {
	return h
}

func main() {
	maxHieght = min(arr[0], arr[len(arr)-1]) // {1,2,3,4}
	waterCaptured := 0

	for index, value := range arr {
		var waterHieght int
		if value > maxHieght {
			value = maxHieght
		}
		if index > 0 && index < len(arr)-1 {
			HieghtLeft := arr[index-1]
			HieghtRight := arr[index+1]
			
			if HieghtLeft > maxHieght && HieghtRight > maxHieght {
				waterHieght = min(arr[index-1], arr[index+1])
			} else {
				waterHieght = maxHieght - value
			}
		}

		waterCaptured += volume(waterHieght)

	}

	fmt.Println(" total water covered = ", waterCaptured)
}
