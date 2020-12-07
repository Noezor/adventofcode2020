package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day05.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

func getPassID(pass string) int {
	return binarySpacePartitionning(stringPassToBinary(pass[:7], "F", "B"))*8 + binarySpacePartitionning(stringPassToBinary(pass[7:], "L", "R"))
}

func binarySpacePartitionning(spacePartitionning []bool) int {
	min := 0
	max := int(math.Pow(float64(2), float64(len(spacePartitionning)))) - 1
	for _, isFront := range spacePartitionning {
		toNewValue := (max - min + 1) / 2
		if isFront {
			max = max - toNewValue
		} else {
			min = min + toNewValue
		}
	}
	return max
}

func stringPassToBinary(pass string, frontChar string, backChar string) []bool {
	arrayBool := make([]bool, len(pass))
	for i, char := range pass {
		if string(char) == frontChar {
			arrayBool[i] = true
		} else if string(char) == backChar {
			arrayBool[i] = false
		} else {
			fmt.Println("UNKNOWN CHAR ", char)
		}
	}
	return arrayBool
}

func max(list []int) int {
	maxValue := 0
	for _, v := range list {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func findMissing(passportIds []int) int {
	sort.Ints(passportIds)
	currentPassportID := passportIds[0]
	for _, passportID := range passportIds[1:] {
		if passportID != (currentPassportID + 1) {
			return currentPassportID + 1
		}
		currentPassportID = passportID
	}
	return -1
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	passportIds := make([]int, len(split))
	for i, s := range split {
		passportIds[i] = getPassID(s)
	}
	fmt.Println(max(passportIds))
	fmt.Println(findMissing(passportIds))
}
