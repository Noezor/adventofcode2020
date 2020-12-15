package main

// template by LFJ

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day15.input", "Relative file path")

func partA(numbers []int) int {
	return playGame(numbers, 2020)
}

func partB(numbers []int) int {
	return playGame(numbers, 30000000)
}

func playGame(numbers []int, nbTurn int) int {
	lastAppearance := make(map[int]int)
	distanceTwoAppearance := make(map[int]int)

	for i, v := range numbers {
		if _, ok := distanceTwoAppearance[v]; !ok {
			distanceTwoAppearance[v] = 0
		} else {
			distanceTwoAppearance[v] = i + 1 - lastAppearance[v]
		}
		lastAppearance[v] = i + 1
	}

	lastSpoken := numbers[len(numbers)-1]
	for turn := len(numbers) + 1; turn <= nbTurn; turn++ {
		if _, ok := distanceTwoAppearance[lastSpoken]; !ok {
			distanceTwoAppearance[lastSpoken] = 0
		} else {
			distanceTwoAppearance[lastSpoken] = turn - 1 - lastAppearance[lastSpoken]
		}
		lastAppearance[lastSpoken] = turn - 1
		lastSpoken = distanceTwoAppearance[lastSpoken]
	}

	return lastSpoken
}

func parseString(contents string) []int {
	split := strings.Split(contents, ",")
	values := make([]int, len(split))
	for i, s := range split {
		value, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return []int{}
		}
		values[i] = value
	}
	return values
}

var testString = "3,1,2"
var expectedOutputA = 1836
var expectedOutputB = 362

func main() {
	flag.Parse()
	contents := testString

	if result := partA(parseString(contents)); result != expectedOutputA {
		fmt.Println("Error, expected", expectedOutputA, "Got ", result)
		return
	}
	if result := partB(parseString(contents)); result != expectedOutputB {
		fmt.Println("Error, expected", expectedOutputB, "Got ", result)
		return
	}
	fmt.Println("Passed tests")

	contents = "9,19,1,6,0,5,4"
	fmt.Println(partA(parseString(contents)))
	fmt.Println(partB(parseString(contents)))
}
