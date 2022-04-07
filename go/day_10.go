package main

import (
	"bufio"
	"fmt"
	. "github.com/pscosta/go-strm/strm"
	"os"
	"sort"
	"strings"
)

var syntax = map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
var illegalPoints = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
var legalPoints = map[string]int{")": 1, "]": 2, "}": 3, ">": 4}

var lines []string

func readInput() {
	file, _ := os.Open("input10.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// parse lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
}

func sol1() int {
	var parsed []string
	totalSum := 0
	for _, line := range lines {
		chars := strings.Split(line, "")
		lineSum := 0
		for _, c := range chars {
			switch {
			case containsKey(syntax, c):
				parsed = append(parsed, c)
			case c != syntax[removeLast(&parsed)]:
				lineSum += illegalPoints[c]
			}
		}
		totalSum += lineSum
		lineSum = 0
	}
	return totalSum
}

func sol2() int {
	var res []int

iterating:
	for _, line := range lines {
		var parsed []string
		chars := strings.Split(line, "")

		for _, c := range chars {
			switch {
			case containsKey(syntax, c):
				parsed = append(parsed, syntax[c])
			case c != removeLast(&parsed):
				continue iterating
			}
		}
		sum := Reduce(From(parsed).Reversed(),
			func(partial int, c string) int { return legalPoints[c] + partial*5 },
		)
		res = append(res, sum)
	}

	sort.Ints(res)
	return res[len(res)/2]
}

func containsKey[K comparable, V any](m map[K]V, key K) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

func removeLast[T any](slice *[]T) (last T) {
	last = (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return
}

func main() {
	readInput()
	fmt.Printf("Sol1: %v\n", sol1())
	fmt.Printf("Sol2: %v\n", sol2())
}
