package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day06.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

func countAtLeastX(dict map[string]int, atLeast int) int {
	count := 0
	for _, value := range dict {
		if value >= atLeast {
			count++
		}
	}
	return count
}

func turnZero(boolDict map[string]int) {
	for key := range boolDict {
		boolDict[key] = 0
	}
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	countAnyYes := 0
	countEveryYes := 0
	questionAnsweredGroupCount := make(map[string]int)
	lenGroup := 0
	for _, s := range split {
		if s == "" {
			countAnyYes += countAtLeastX(questionAnsweredGroupCount, 1)
			countEveryYes += countAtLeastX(questionAnsweredGroupCount, lenGroup)
			turnZero(questionAnsweredGroupCount)
			lenGroup = 0
		} else {
			lenGroup++
			for _, char := range s {
				if _, ok := questionAnsweredGroupCount[string(char)]; !ok {
					questionAnsweredGroupCount[string(char)] = 1
				} else {
					questionAnsweredGroupCount[string(char)] = questionAnsweredGroupCount[string(char)] + 1
				}
			}
		}
	}
	countAnyYes += countAtLeastX(questionAnsweredGroupCount, 1)
	countEveryYes += countAtLeastX(questionAnsweredGroupCount, lenGroup)
	fmt.Println(countAnyYes)
	fmt.Println(countEveryYes)
}
