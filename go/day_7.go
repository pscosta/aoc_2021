package main

import (
	"bufio"
	"fmt"
	. "github.com/pscosta/go-strm/strm"
	"math"
	"os"
	"strconv"
	"strings"
)

var crabs []int

func readInput() {
	file, _ := os.Open("input7.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// parse lines
	for scanner.Scan() {
		intCrabs := Map(From(strings.Split(scanner.Text(), ",")),
			func(f string) int { fi, _ := strconv.Atoi(f); return fi },
		).ToSlice()
		crabs = RangeFrom(intCrabs).Sorted().ToSlice()
	}
}

func main() {
	readInput()
	var avg = Sum(From(crabs)) / len(crabs)
	var middle = crabs[len(crabs)/2]

	fmt.Printf("Sol1: %v\n", From(crabs).SumBy(
		func(it int) int { return int(math.Abs(float64(middle - it))) }),
	)

	fmt.Printf("Sol2: %v\n", From(crabs).SumBy(
		func(it int) int {
			return int((math.Abs(float64(avg-it))*math.Abs(float64(avg-it)) + math.Abs(float64(avg-it))) / 2)
		}))
}
