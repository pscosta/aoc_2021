package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	dir    string
	amount int
}

func readInput() (instructions []Input) {
	file, _ := os.Open("input2.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		var inst Input
		fields := strings.Fields(scanner.Text())
		inst.dir = fields[0]
		inst.amount, _ = strconv.Atoi(fields[1])
		instructions = append(instructions, inst)
	}
	return
}

var input = readInput()

func main() {
	fmt.Printf("sol1: %v\n\n", sol1())
	fmt.Printf("sol2: %v\n\n", sol2())
}

func sol1() (res int) {
	horiz := 0
	depth := 0

	for _, inst := range input {
		switch inst.dir {
		case "forward":
			horiz += inst.amount
		case "down":
			depth += inst.amount
		case "up":
			depth -= inst.amount
		}
	}
	return horiz * depth
}

func sol2() (res int) {
	horiz := 0
	depth := 0
	aim := 0

	for _, inst := range input {
		switch inst.dir {
		case "forward":
			horiz += inst.amount
			depth += aim * inst.amount
		case "down":
			aim += inst.amount
		case "up":
			aim -= inst.amount
		}
	}
	return horiz * depth
}
