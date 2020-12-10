package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day10.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

func plugAll(adapters []int) []int {
	// returns list of plugable adapters in an order which is plugable
	deviceAdapter := getDeviceAdapter(adapters)
	allAdapters := append(adapters, deviceAdapter)
	sort.Ints(allAdapters)
	return allAdapters
}

func getDeviceAdapter(adapters []int) int {
	return max(adapters) + 3
}

func max(numbers []int) int {
	max := numbers[0]
	for _, v := range numbers[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func getDifferences(plugableList []int) map[int]int {
	currentAdapter := 0
	differences := make(map[int]int)
	for _, adapter := range plugableList {
		difference := adapter - currentAdapter
		currentAdapter = adapter
		differences[difference] = differences[difference] + 1
	}
	return differences
}

func getNbCombinaisons(adapters []int) int {

	sortedAdapters := adapters
	sort.Ints(sortedAdapters)
	nbCombinaisonsStartsWith := make(map[int]int, len(sortedAdapters))
	nbCombinaisonsStartsWith[sortedAdapters[len(sortedAdapters)-1]] = 1
	for i := len(sortedAdapters) - 2; i >= 0; i-- {
		valueStartWith := sortedAdapters[i]
		nbCombinaisonsStartsWith[valueStartWith] = nbCombinaisonsStartsWith[valueStartWith+1] + nbCombinaisonsStartsWith[valueStartWith+2] + nbCombinaisonsStartsWith[valueStartWith+3]
	}
	return nbCombinaisonsStartsWith[1] + nbCombinaisonsStartsWith[2] + nbCombinaisonsStartsWith[3]
}

func getKey(differences map[int]int) int {
	return differences[3] * differences[1]
}

func isEqual(map1, map2 map[int]int) bool {
	for k := range map1 {
		if map1[k] != map2[k] {
			return false
		}
	}
	return true
}

var TestString = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

var expectedDifference = map[int]int{1: 22, 3: 10}
var expectedCombinaisons = 19208

func main() {
	contents := TestString
	split := strings.Split(contents, "\n")
	adapters := make([]int, len(split))
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Failed to parse ", s)
			return
		}
		adapters[i] = n
	}
	if !isEqual(getDifferences(plugAll(adapters)), expectedDifference) {
		fmt.Println("Obtained ", getDifferences(plugAll(adapters)), " expected", expectedDifference)
		return
	}
	if getNbCombinaisons(adapters) != expectedCombinaisons {
		fmt.Println("Obtained ", getNbCombinaisons(adapters), " expected", expectedCombinaisons)
		return
	}

	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents = string(bytes)
	split = strings.Split(contents, "\n")
	adapters = make([]int, len(split))
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		adapters[i] = n
	}
	fmt.Println(getKey(getDifferences(plugAll(adapters))))
	fmt.Println(getNbCombinaisons(adapters))
}
