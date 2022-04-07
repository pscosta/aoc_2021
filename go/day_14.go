package main

import (
	"bufio"
	"fmt"
	. "github.com/pscosta/go-strm/strm"
	"math"
	"os"
	"strings"
)

var rules = make(map[string]string)
var template []string

func readInput() {
	file, _ := os.Open("input14.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	scanner.Scan()
	template = strings.Split(scanner.Text(), "")
	scanner.Scan()

	// parse lines
	for scanner.Scan() {
		rule := strings.Split(scanner.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}
}

func sol1(iterations int) int {
	var newTemplate = make([]string, len(template))
	copy(newTemplate, template)

	for iter := 0; iter < iterations; iter++ {
		windows := From(newTemplate).Windowed(2, 1)
		for i, window := range windows {
			newChar := rules[window[0]+window[1]]
			windows[i] = append(window[:1], append([]string{newChar}, window[1:]...)...)
		}

		newTemplate = make([]string, 0, len(windows))
		for _, window := range windows {
			newTemplate = append(newTemplate, window[0:2]...)
		}
		newTemplate = append(newTemplate, windows[len(windows)-1][2])
	}

	min := math.MaxInt
	max := 0
	for _, v := range GroupBy(From(newTemplate), func(it string) string { return it }) {
		if len(v) < min {
			min = len(v)
		}
		if len(v) > max {
			max = len(v)
		}
	}
	return max - min
}

func sol2(iterations int) int {
	pairs := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		pairs[template[i]+template[i+1]]++
	}

	for iter := 0; iter < iterations; iter++ {
		updatedPairs := make(map[string]int)
		for k, v := range pairs {
			updatedPairs[string(k[0])+rules[k]] += v
			updatedPairs[rules[k]+string(k[1])] += v
		}
		pairs = updatedPairs
	}

	sums := make(map[string]int)
	for k, v := range pairs {
		sums[string(k[0])] += v
	}
	max := 0
	min := math.MaxInt
	for _, v := range sums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min + 1
}

func main() {
	readInput()
	fmt.Printf("Sol1: %v\n", sol1(10))
	fmt.Printf("Sol2: %v\n", sol2(40))
}
