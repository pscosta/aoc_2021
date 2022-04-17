package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

var drawNumbers []int
var boards [][][]int

func readInput() {
	file, _ := os.Open("input4.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// parse draw numbers
	scanner.Scan()
	drawNumbers = toInts(s.Split(scanner.Text(), ","))

	// parse remaining file
	var board [][]int
	for scanner.Scan() {
		row := toInts(s.Split(s.TrimSpace(s.ReplaceAll(scanner.Text(), "  ", " ")), " "))
		if len(row) > 1 {
			board = append(board, row)
		}
		if len(board) == 5 {
			boards = append(boards, board)
			board = make([][]int, 0)
		}
	}
}

func toInts(tokens []string) (res []int) {
	for _, token := range tokens {
		number, _ := strconv.Atoi(token)
		res = append(res, number)
	}
	return
}

func isWinner(board [][]int, drawnNums []int) bool {
	for _, row := range board {
		markedCount := 0
		hash := make(map[int]bool)
		for _, e := range drawnNums {
			hash[e] = true
		}
		for _, e := range row {
			if hash[e] {
				markedCount++
			}
		}
		if len(row) == markedCount {
			return true
		}
	}
	return false
}

func reverseBoard(board [][]int) (reversedBoard [][]int) {
	reversedBoard = make([][]int, len(board))
	for _, row := range board {
		for i := range row {
			reversedBoard[i] = append(reversedBoard[i], row[i])
		}
	}
	return
}

func Score(board [][]int, drawnNums []int, lastDrawn int) (score int) {
	for _, row := range board {
		hash := make(map[int]bool)
		for _, e := range drawnNums {
			hash[e] = true
		}
		for _, e := range row {
			if !hash[e] {
				score += e
			}
		}
	}
	return score * lastDrawn
}

func sol1() (score int) {
	var drawnSoFar []int
	for _, drawnNum := range drawNumbers {
		drawnSoFar = append(drawnSoFar, drawnNum)

		for _, board := range boards {
			if isWinner(board, drawnSoFar) || isWinner(reverseBoard(board), drawnSoFar) {
				return Score(board, drawnSoFar, drawnNum)
			}
		}
	}
	return -1
}

func sol2() (score int) {
	var drawnSoFar []int
	winners := make(map[int]bool)

	for _, drawnNum := range drawNumbers {
		drawnSoFar = append(drawnSoFar, drawnNum)

		for idx, board := range boards {
			if !winners[idx] && (isWinner(board, drawnSoFar) || isWinner(reverseBoard(board), drawnSoFar)) {
				winners[idx] = true
			}
			if len(winners) == len(boards) {
				return Score(board, drawnSoFar, drawnNum)
			}
		}
	}
	return -1
}

func main() {
	readInput()
	fmt.Printf("sol1: %v\n", sol1())
	fmt.Printf("sol2: %v\n", sol2())
}
