package main

import (
	"bufio"
	"fmt"
	. "github.com/pscosta/go-strm/strm"
	"os"
	"strconv"
	"strings"
)

var initialFish []int

func readInput() {
	file, _ := os.Open("input6.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// parse lines
	for scanner.Scan() {
		initialFish = Map(From(strings.Split(scanner.Text(), ",")),
			func(f string) int { fi, _ := strconv.Atoi(f); return fi },
		).ToSlice()
	}
}

func values[K comparable, V any](m map[K]V) (values []V) {
	for _, v := range m {
		values = append(values, v)
	}
	return
}

func lanternFish(days int) int {
	allSortedFish := GroupBy(CopyFrom(initialFish), func(it int) int { return it })
	newFish := make(map[int]int)
	allFish := make(map[int]int)
	for k, v := range allSortedFish {
		allFish[k] = len(v)
	}

	for i := 0; i < days; i++ {
		for age, fishCount := range allFish {
			switch age {
			case 0:
				newFish[6] += fishCount
			default:
				newFish[age-1] += fishCount
			}
		}
		// update new fish count
		if zeroFishCount, ok := allFish[0]; ok {
			newFish[8] = zeroFishCount
		}
		// update all fish with new ages
		allFish = make(map[int]int)
		for k, v := range newFish {
			allFish[k] += v
		}
		// clear newFish for next iteration
		newFish = make(map[int]int)
	}
	return Sum(From(values(allFish)))
}

func main() {
	readInput()
	fmt.Printf("sol1: %v\n", lanternFish(80))
	fmt.Printf("sol1: %v\n", lanternFish(256))
}
