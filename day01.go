package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path")
var partB = flag.Bool("partB", true, "Using part B logic")

var targetSum int = 2020

func findProduct3ToSum(numbers []int, targetSum int) (int, int, int) {
	for _, n := range numbers {
		b, c := findProduct2ToSum(numbers, targetSum-n)
		if (b != 0) && (c != 0) {
			return n, b, c
		}
	}
	return 0, 0, 0
}

func findProduct2ToSum(numbers []int, targetSum int) (int, int) {
	hashmapComplements := buildHashmapComplement(numbers, targetSum)
	a, b := findPair(numbers, hashmapComplements, targetSum)
	return a, b
}

func buildHashmapComplement(numbers []int, targetSum int) map[int]bool {
	hashmapComplements := make(map[int]bool)
	for _, n := range numbers {
		hashmapComplements[targetSum-n] = true
	}
	return hashmapComplements
}

func findPair(numbers []int, hashmapComplements map[int]bool, targetSum int) (int, int) {

	for _, n := range numbers {
		if hashmapComplements[n] {
			return n, targetSum - n
		}
	}
	return 0, 0
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	numbers := make([]int, len(split))
	for i, s := range split {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		numbers[i] = n
	}
	if *partB {
		a, b, c := findProduct3ToSum(numbers, targetSum)

		fmt.Println("Target product VARIATION B:")
		fmt.Println("%n", a)
		fmt.Println("%n", b)
		fmt.Println("%n", c)
	} else {
		a, b := findProduct2ToSum(numbers, targetSum)

		fmt.Println("Target product :")
		fmt.Println("%n", a)
		fmt.Println("%n", b)
	}
}
