package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() (measurements []int) {
	file, _ := os.Open("input1.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, val)
	}
	return
}

var input = readInput()

func main() {
	fmt.Printf("sol1: %v\n", sol1())
	fmt.Printf("sol2: %v\n", sol2())
}

func sol1() (res int) {
	prev := input[0]

	for _, m := range input {
		if m > prev {
			res++
		}
		prev = m
	}
	return
}

func sol2() (res int) {
	prev := input[0] + input[1] + input[2]

	for i := 0; i <= len(input)-3; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		if sum > prev {
			res++
		}
		prev = sum
	}
	return
}
