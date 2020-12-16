package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day13.input", "Relative file path")

func partA(firstTimestamp int, busIds []int) int {
	minTimestamp := 1000000000
	minId := -1

	for _, busId := range busIds {
		if closestTimestampBus := closestBusToTimestamp(firstTimestamp, busId); closestTimestampBus < minTimestamp {
			minTimestamp = closestTimestampBus
			minId = busId
		}
	}
	return minId * (minTimestamp - firstTimestamp)
}

func partB(busIds, offsets []int) int {
	n, x := chineseReminderSolution(offsets, busIds)
	for x < 0 {
		x += n
	}
	return x
}

func bruteforceChineseSolution(startingValue int, increment int, busIds []int, offsets []int) int {
	currentValue := startingValue
	for !verifiesChineseReminder(currentValue, offsets, busIds) {
		currentValue += increment
	}
	return currentValue
}

func closestBusToTimestamp(timestamp int, busId int) int {
	timestampBusDeparture := busId
	for timestamp > timestampBusDeparture {
		timestampBusDeparture += busId
	}
	return timestampBusDeparture
}

func verifiesChineseReminder(x int, a []int, n []int) bool {
	if len(a) != len(n) {
		fmt.Println("Wrong input size", a, n)
	}
	for i, ni := range n {
		ai := a[i]
		mod := x % ni
		if mod < 0 {
			mod = mod + ni
		}
		if mod != ai {
			return false
		}
	}
	return true
}

func chineseReminderSolution(a, n []int) (int, int) {
	if len(a) != len(n) {
		fmt.Println("input of wrong size", a, n)
		return 0, -1
	}
	aCurrent, nCurrent := a[0], n[0]
	for i := 1; i < len(a); i++ {
		x := chineseReminderSolutionPair(aCurrent, nCurrent, a[i], n[i])
		nCurrent = nCurrent * n[i]
		aCurrent = x % nCurrent
	}
	return nCurrent, aCurrent
}

func chineseReminderSolutionPair(a1, n1, a2, n2 int) int {
	m1, m2 := bezoutCoefficients(n1, n2)
	if n1*m1+n2*m2 != 1 {
		fmt.Println("Error, n1*m1 + n2*m2 == ", n1*m1+n2*m2, n1, m1, n2, m2)
	}
	return a2*n1*m1 + a1*m2*n2
}

func bezoutCoefficients(a, b int) (int, int) {
	old_r, r := a, b
	old_s, s := 1, 0
	old_t, t := 0, 1

	for r != 0 {
		quotient := old_r / r
		old_r, r = r, old_r-quotient*r
		old_s, s = s, old_s-quotient*s
		old_t, t = t, old_t-quotient*t
	}

	return old_s, old_t
}

func parseStringA(contents string) (int, []int) {
	split := strings.Split(contents, "\n")
	firstTimestamp, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("Failed to parse", split[0])
		return 0, []int{}
	}

	busIds := make([]int, 0)
	splitIds := strings.Split(split[1], ",")
	for _, s := range splitIds {
		value, err := strconv.Atoi(s)
		if err == nil {
			busIds = append(busIds, value)
		}
	}

	return firstTimestamp, busIds
}

func parseStringB(contents string) ([]int, []int) {
	split := strings.Split(contents, "\n")

	busIds := make([]int, 0)
	offsets := make([]int, 0)
	splitIds := strings.Split(split[1], ",")
	for i, s := range splitIds {
		busId, err := strconv.Atoi(s)
		if err == nil {
			busIds = append(busIds, busId)
			offset := (busId - i) % busId
			for offset < 0 {
				offset += busId
			}
			offsets = append(offsets, offset)
		}
	}
	return busIds, offsets
}

var testString = `939
7,13,x,x,59,x,31,19`
var expectedOutput = 295
var expectedOutputB = 1068781

func main() {
	flag.Parse()
	contents := testString

	if result := partA(parseStringA(contents)); result != expectedOutput {
		fmt.Println("Error, expected", expectedOutput, "Got ", result)
		return
	}
	if result := partB(parseStringB(contents)); result != expectedOutputB {
		fmt.Println("Error, expected", expectedOutputB, "Got ", result)
		return
	}
	fmt.Println("Passed tests")

	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents = string(bytes)
	fmt.Println(partA(parseStringA(contents)))

	fmt.Println("partB with constructive algorithm", partB(parseStringB(contents)))
	// magic values from the result of chinese reminder for the first 6 elements :
	busIds, offsets := parseStringB(contents)
	fmt.Println("Correct solution ", bruteforceChineseSolution(4241072787, 5010216523, busIds, offsets))
}
