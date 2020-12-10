package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day09.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

func findFirstNotSumOfTwo(numbers []int, preambleSize int) int {
	for i := preambleSize; i < len(numbers); i++ {
		if !isSumOfTwoBruteforce(numbers[i], numbers[(i-preambleSize):i]) {
			return numbers[i]
		}
	}
	return -1
}

func isSumOfTwoBruteforce(target int, numbers []int) bool {
	for _, a := range numbers {
		for _, b := range numbers {
			if a+b == target && a != b {
				return true
			}
		}
	}
	return false
}

func parseFileString(contents string) []int {
	split := strings.Split(contents, "\n")

	numbers := make([]int, len(split))
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		numbers[i] = n
	}
	return numbers
}

func findContinuousSumsTo(numbers []int, target int) []int {
	for continuousSize := 2; continuousSize < len(numbers); continuousSize++ {
		solution := findContinuousSumToAtGivenSize(numbers, continuousSize, target)
		if solution != nil {
			fmt.Println("continuous sums to solution :", solution, "target :", target)
			return solution
		}
	}
	return nil
}

func findContinuousSumToAtGivenSize(numbers []int, continuousSize int, target int) []int {
	currentSum := sum(numbers[:continuousSize])
	if currentSum == target {
		return numbers[:continuousSize]
	}
	fmt.Println("first sum", currentSum)

	for beginningSum := 1; beginningSum < len(numbers)-continuousSize; beginningSum++ {
		currentSum = currentSum + numbers[beginningSum+continuousSize-1] - numbers[beginningSum-1]
		fmt.Println(beginningSum, continuousSize, target, currentSum)
		if currentSum == target {
			return numbers[beginningSum:(beginningSum + continuousSize)]
		}
	}
	return nil
}

func sum(numbers []int) int {
	sumNumbers := 0
	for _, v := range numbers {
		sumNumbers += v
	}
	return sumNumbers
}

func getEncryptionWeakness(continuousArraySum []int) int {
	return min(continuousArraySum) + max(continuousArraySum)
}

func min(numbers []int) int {
	minValue := numbers[0]
	for _, v := range numbers[1:] {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

func max(numbers []int) int {
	maxValue := numbers[0]
	for _, v := range numbers[1:] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

var testContent string = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func main() {
	contents := testContent
	wrongNumber := 127
	numbers := parseFileString(contents)
	if findFirstNotSumOfTwo(numbers, 5) != wrongNumber {
		fmt.Println(numbers)
		fmt.Println("Error, expected", wrongNumber, "but found ", findFirstNotSumOfTwo(numbers, 5), "on test sample")
		return
	}
	expectedWeakness := 62
	solution := getEncryptionWeakness(findContinuousSumsTo(numbers, wrongNumber))
	if solution != expectedWeakness {
		fmt.Println(numbers)
		fmt.Println("Error, expected", expectedWeakness, "but found ", solution, "on test sample")
		return
	}

	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents = string(bytes)
	numbers = parseFileString(contents)
	wrongNumber = findFirstNotSumOfTwo(numbers, 25)
	fmt.Println(getEncryptionWeakness(findContinuousSumsTo(numbers, wrongNumber)))
}
