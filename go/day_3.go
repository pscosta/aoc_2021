package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() (input [][]rune) {
	file, _ := os.Open("input3.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		binaryChars := []rune(scanner.Text())
		input = append(input, binaryChars)
	}
	return
}

func reverseInput() [][]rune {
	input := readInput()
	reversedInput := make([][]rune, len(input[0]))

	for _, binary := range input {
		for j := 0; j < len(binary); j++ {
			reversedInput[j] = append(reversedInput[j], binary[j])
		}
	}
	return reversedInput
}

func countBinary(input []*[]rune, i int) (zeros int, ones int) {
	for _, runes := range input {
		zeros += strings.Count(string((*runes)[i]), "0")
		ones += strings.Count(string((*runes)[i]), "1")
	}
	return zeros, ones
}

func sol1() int64 {
	gamaBin := ""
	epsilonBin := ""

	for _, runes := range reverseInput() {
		zeros := strings.Count(string(runes), "0")
		ones := strings.Count(string(runes), "1")
		if zeros > ones {
			gamaBin += "0"
			epsilonBin += "1"
		} else {
			gamaBin += "1"
			epsilonBin += "0"
		}
	}

	gama, _ := strconv.ParseInt(gamaBin, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonBin, 2, 64)
	return gama * epsilon
}

func sol2() int64 {
	input := readInput()
	o2Input := make([]*[]rune, len(input))
	co2Input := make([]*[]rune, len(input))

	for idx := range input {
		o2Input[idx] = &(input[idx])
		co2Input[idx] = &(input[idx])
	}

	for i := 0; i < len(input[0]); i++ {
		zeros, ones := countBinary(o2Input, i)

		for idx := len(o2Input) - 1; idx >= 0; idx-- {
			digit := string((*(o2Input[idx]))[i])

			if zeros > ones {
				if digit == "1" && len(o2Input) > 1 {
					remove(&o2Input, idx)
				}
			} else {
				if digit == "0" && len(o2Input) > 1 {
					remove(&o2Input, idx)
				}
			}
		}
	}

	for i := 0; i < len(input[0]); i++ {
		zeros, ones := countBinary(co2Input, i)

		for idx := len(co2Input) - 1; idx >= 0; idx-- {
			binary := string((*(co2Input[idx]))[i])

			if zeros > ones {
				if binary == "0" && len(co2Input) > 1 {
					co2Input = append(co2Input[:idx], co2Input[idx+1:]...)
				}
			} else {
				if binary == "1" && len(co2Input) > 1 {
					co2Input = append(co2Input[:idx], co2Input[idx+1:]...)
				}
			}
		}
	}

	o2, _ := strconv.ParseInt(string(*o2Input[0]), 2, 64)
	co2, _ := strconv.ParseInt(string(*co2Input[0]), 2, 64)
	return o2 * co2
}

func remove(input *[]*[]rune, idx int) {
	*input = append((*input)[:idx], (*input)[idx+1:]...)
}

func main() {
	fmt.Printf("sol1: %v\n", sol1())
	fmt.Printf("sol2: %v\n", sol2())
}
