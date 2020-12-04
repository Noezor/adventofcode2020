package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day03.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

func nbTrees(grid []string, movesRight int, movesDown int) int {
	tokenTree := "#"
	nbTrees := 0

	currentX := 0
	currentY := 0
	for currentY < len(grid) {
		currentRow := grid[currentY%len(grid)]
		if string(currentRow[currentX%len(currentRow)]) == tokenTree {
			nbTrees++
		}
		currentX += movesRight
		currentY += movesDown
	}
	return nbTrees
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	grid := make([]string, len(split))
	for i, s := range split {
		if s == "" {
			continue
		}
		grid[i] = s
	}
	fmt.Println("%n/n", nbTrees(grid, 1, 1)*nbTrees(grid, 3, 1)*nbTrees(grid, 5, 1)*nbTrees(grid, 7, 1)*nbTrees(grid, 1, 2))
}
