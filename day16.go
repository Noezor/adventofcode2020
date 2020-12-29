package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day16.input", "Relative file path")

type rule struct {
	name      string
	intervals [2]interval
}

type interval struct {
	begin int
	end   int
}

type ticket struct {
	numbers []int
}

func partA(rules []rule, myTicket ticket, nearbyTickets []ticket) int {
	errorRate := 0
	for _, t := range nearbyTickets {
		for _, n := range t.numbers {
			matchesOneRule := false
			for _, r := range rules {
				if r.matches(n) {
					matchesOneRule = true
				}
			}
			if !matchesOneRule {
				errorRate += n
			}
		}
	}
	return errorRate
}

func partB(rules []rule, myTicket ticket, nearbyTickets []ticket) int {
	correspondences := solveCorrespondences(rules, myTicket, nearbyTickets)
	return scoreTicket(myTicket, correspondences)
}

func solveCorrespondences(rules []rule, myTicket ticket, nearbyTickets []ticket) map[string]int {
	validTickets := make([]ticket, 0)
	for _, t := range nearbyTickets {
		if t.matchesAtLeastOne(rules) {
			validTickets = append(validTickets, t)
		}
	}

	possibleRulesFields := make([][]string, len(rules))

	for iField := 0; iField < len(rules); iField++ {
		numbers := make([]int, len(validTickets))
		for i, t := range validTickets {
			numbers[i] = t.numbers[iField]
		}
		possibleRulesFields[iField] = possibleRules(numbers, rules)
	}
	return solvePossibleRules(possibleRulesFields, rules)
}

func solvePossibleRules(possibleRulesFields [][]string, rules []rule) map[string]int {
	found := make(map[string]bool, len(rules))
	correspondences := make(map[string]int, 0)

	for len(correspondences) != len(rules) {
		for i, rulesField := range possibleRulesFields {
			possibleField := areIn(rulesField, found)
			if len(possibleField) == 1 {
				correspondingRule := possibleField[0]
				found[correspondingRule] = true
				correspondences[correspondingRule] = i
			}
		}
	}
	return correspondences
}

func scoreTicket(t ticket, correspondences map[string]int) int {
	score := 1
	for ruleName, i := range correspondences {
		if strings.HasPrefix(ruleName, "departure") {
			score *= t.numbers[i]
		}
	}
	return score
}

func areIn(s []string, m map[string]bool) []string {
	filtered := make([]string, 0)

	for _, c := range s {
		if !m[c] {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func possibleRules(numbers []int, rules []rule) []string {
	possible := make(map[string]bool, len(rules))
	for _, r := range rules {
		possible[r.name] = true
	}
	for _, r := range rules {
		for _, n := range numbers {
			if !r.matches(n) {
				possible[r.name] = false
			}
		}
	}

	output := make([]string, 0)
	for _, r := range rules {
		if possible[r.name] {
			output = append(output, r.name)
		}
	}
	return output
}

func (t ticket) matchesAtLeastOne(rules []rule) bool {
	isValid := true
	for _, n := range t.numbers {
		matchesOneRule := false
		for _, r := range rules {
			if r.matches(n) {
				matchesOneRule = true
			}
		}
		if !matchesOneRule {
			isValid = false
		}
	}
	return isValid
}

func (r rule) matches(x int) bool {
	return (r.intervals[0].contains(x) || r.intervals[1].contains(x))
}

func (i interval) contains(x int) bool {
	return (i.begin <= x && x <= i.end)
}

func parseString(contents string) ([]rule, ticket, []ticket) {
	split := strings.Split(contents, "\n")
	rules := make([]rule, 0)
	var myTicket ticket
	nearbyTickets := make([]ticket, 0)

	for i, s := range split {
		if s == "your ticket:" {
			rules = parseRules(split[:i-1])
			myTicket = parseTicket(split[i+1])
		}
		if s == "nearby tickets:" {
			for j := i + 1; j < len(split); j++ {
				nearbyTickets = append(nearbyTickets, parseTicket(split[j]))
			}
		}
	}
	return rules, myTicket, nearbyTickets
}

func parseRules(split []string) []rule {
	rules := make([]rule, len(split))
	for i, s := range split {
		rules[i] = parseRule(s)
	}
	return rules
}

func parseRule(s string) rule {
	split := strings.Split(s, ":")
	name := split[0]
	splitIntervals := strings.Split(split[1], " ")[1:]
	intervals := [2]interval{parseInterval(splitIntervals[0]), parseInterval(splitIntervals[len(splitIntervals)-1])}
	return rule{name: name, intervals: intervals}
}

func parseInterval(s string) interval {
	split := strings.Split(s, "-")
	beg, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to parse ", split[0])
	}
	end, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to parse ", split[1])
	}
	return interval{begin: beg, end: end}
}

func parseTicket(s string) ticket {
	split := strings.Split(s, ",")
	numbers := make([]int, len(split))
	for i, c := range split {
		n, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Failed to parse", c)
		}
		numbers[i] = n
	}
	return ticket{numbers: numbers}
}

var testString = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
var expectedOutput = 71

func main() {
	flag.Parse()
	contents := testString

	if result := partA(parseString(contents)); result != expectedOutput {
		fmt.Println("Error, expected", expectedOutput, "Got ", result)
		return
	}
	fmt.Println(partB(parseString(contents)))
	fmt.Println("Passed tests")

	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Failed to parse file")
		return
	}
	contents = string(bytes)
	fmt.Println(partA(parseString(contents)))
	fmt.Println(partB(parseString(contents)))
}
