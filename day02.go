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
var partB = flag.Bool("partB", true, "Using part B logic")

type PolicyA struct {
	minOccurence int
	maxOccurence int
	char         string
}

type PolicyB struct {
	firstOccurence int
	lastOccurence  int
	char           string
}

func isValidLineA(inputLine string) bool {
	policy, password := splitInputLineA(inputLine)
	return isValidA(password, policy)
}

func isValidLineB(inputLine string) bool {
	policy, password := splitInputLineB(inputLine)
	return isValidB(password, policy)
}

func splitInputLineA(inputLine string) (PolicyA, string) {
	// 2-9 c: ccccccccc
	splitSpace := strings.Split(inputLine, " ")
	min, max := splitMinMax(splitSpace[0])
	character := string(splitSpace[1][0])
	password := splitSpace[2]
	return PolicyA{minOccurence: min, maxOccurence: max, char: character}, password
}

func splitInputLineB(inputLine string) (PolicyB, string) {
	// 2-9 c: ccccccccc
	splitSpace := strings.Split(inputLine, " ")
	min, max := splitMinMax(splitSpace[0])
	character := string(splitSpace[1][0])
	password := splitSpace[2]
	return PolicyB{firstOccurence: min, lastOccurence: max, char: character}, password
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

func isValidA(password string, policyPassword PolicyA) bool {
	policyCharCount := strings.Count(password, policyPassword.char)
	return (policyPassword.minOccurence <= policyCharCount) && (policyCharCount <= policyPassword.maxOccurence)
}

func isValidB(password string, policyPassword PolicyB) bool {
	isAsFirst := isOccurence(password, policyPassword.firstOccurence, policyPassword.char)
	isAsLast := isOccurence(password, policyPassword.lastOccurence, policyPassword.char)
	return (isAsFirst != isAsLast)
}

func isOccurence(password string, index int, character string) bool {
	if 0 <= index-1 && index-1 < len(password) {
		return string(password[index-1]) == character
	} else {
		return false
	}
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

	if *partB {
		for _, s := range split {
			if isValidLineB(s) {
				sum++
			}
		}
	} else {
		for _, s := range split {
			if isValidLineA(s) {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
