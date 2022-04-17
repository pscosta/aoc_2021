package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

type Line struct {
	x1, x2, y1, y2 int
}

var lines []Line
var matrix [][]int

func readInput() {
	file, _ := os.Open("input5.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// parse lines
	for scanner.Scan() {
		coords := toInts(s.Split(s.ReplaceAll(scanner.Text(), " ", ""), "->"))
		lines = append(lines, Line{x1: coords[0], y1: coords[1], x2: coords[2], y2: coords[3]})
	}
}

func initMatrix() {
	maxX := 0
	maxY := 0
	for _, line := range lines {
		maxX = maxOf(maxX, line.x1, line.x2)
		maxY = maxOf(maxY, line.y1, line.y2)
	}

	matrix = make([][]int, maxY+1)
	for i := range matrix {
		matrix[i] = make([]int, maxX+1)
	}
}

func fillMatrix(countDiagonals bool) {
	for _, line := range lines {
		x1 := minOf(line.x1, line.x2)
		x2 := maxOf(line.x1, line.x2)
		y1 := minOf(line.y1, line.y2)
		y2 := maxOf(line.y1, line.y2)

		if x1 == x2 {
			for y := y1; y <= y2; y++ {
				matrix[y][x1] += 1
			}
		} else if y1 == y2 {
			for x := x1; x <= x2; x++ {
				matrix[y1][x] += 1
			}
		} else if countDiagonals {
			x := line.x1
			y := line.y1
			for i := 0; i <= absInt(x1, x2); i++ {
				matrix[y][x] += 1
				if line.x2 > x {
					x++
				} else {
					x--
				}
				if line.y2 > y {
					y++
				} else {
					y--
				}
			}
		}
	}
}

func absInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func countIntersections() (intersections int) {
	for _, line := range matrix {
		for _, num := range line {
			if num > 1 {
				intersections++
			}
		}
	}
	return
}

func maxOf(nums ...int) (max int) {
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}

func minOf(nums ...int) (min int) {
	min = nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return
}

func toInts(pairs []string) (res []int) {
	for _, pair := range pairs {
		for _, token := range s.Split(pair, ",") {
			number, _ := strconv.Atoi(token)
			res = append(res, number)
		}
	}
	return
}

func sol1() int {
	initMatrix()
	fillMatrix(false)
	return countIntersections()
}

func sol2() int {
	initMatrix()
	fillMatrix(true)
	return countIntersections()
}

func main() {
	readInput()
	fmt.Printf("sol1: %v\n", sol1())
	fmt.Printf("sol2: %v\n", sol2())
}
