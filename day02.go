package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

type Policy struct {
	minOccurence int
	maxOccurence int
	char         string
}

func isValidLine(inputLine string) bool {
	policy, password := splitInputLine(inputLine)
	return isValid(password, policy)
}

func splitInputLine(inputLine string) (Policy, string) {
	// 2-9 c: ccccccccc
	splitSpace := strings.Split(inputLine, " ")
	min, max := splitMinMax(splitSpace[0])
	character := string(splitSpace[1][0])
	password := splitSpace[2]
	return Policy{minOccurence: min, maxOccurence: max, char: character}, password
}

func splitMinMax(minMax string) (int, int) {
	splitDash := strings.Split(minMax, "-")
	min, err := strconv.Atoi(splitDash[0])
	if err != nil {
		fmt.Printf("Failed to parse %s", splitDash[0])
	}
	max, err := strconv.Atoi(splitDash[1])
	if err != nil {
		fmt.Printf("Failed to parse %s", splitDash[1])
	}
	return min, max
}

func isValid(password string, policyPassword Policy) bool {
	policyCharCount := strings.Count(password, policyPassword.char)
	return (policyPassword.minOccurence <= policyCharCount) && (policyCharCount <= policyPassword.maxOccurence)
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	sum := 0
	split := strings.Split(contents, "\n")
	for _, s := range split {
		if isValidLine(s) {
			sum++
		}
	}
	fmt.Println(sum)
}
