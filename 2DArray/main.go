package main

import "fmt"

var chessBoard [8][8]string

// ## 2. Given a Chessboard and a File, count how many squares are occupied

func checkOccupiedRank(rank int) int {
	count := 0
	for _, value := range chessBoard[rank] {
		if value == "#" {
			count++
		}
	}
	return count
}

func checkOccupiedFiles(file int) int {
	count := 0

	for salmanKhan, _ := range chessBoard {
		if chessBoard[salmanKhan][file] == "#" {
			count++
		}
	}

	return count
}

func CountAll() int {
	count := 0
	for index, _ := range chessBoard {
		for i := 0; i < len(chessBoard[index]); i++ {
			count++
		}
	}
	return count
}

func countOccupied() int {
	count := 0
	for index, _ := range chessBoard {
		for _, value := range chessBoard[index] {
			if value == "#" {
				count++
			}
		}
	}
	/*
		for i:=0; i<len(chessBoard); i++{
			for j:=0 ; j<len(chessBoard[i]) ; j++{
				if chessBoard[i][j] == "#"{
					count++
				}
			}
		}
	*/
	return count
}

func main() {
	chessBoard = [8][8]string{
		{"#", "_", "#", "#", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"#", "#", "_", "_", "#", "#", "#", "#"},
		{"_", "_", "#", "#", "_", "_", "_", "#"},
		{"_", "_", "_", "_", "#", "_", "_", "#"},
		{"#", "_", "_", "_", "_", "_", "_", "#"},
		{"_", "#", "#", "#", "_", "_", "_", "#"},
		{"#", "_", "_", "_", "_", "#", "_", "#"},
	}

	fmt.Println(checkOccupiedFiles(2))

}
