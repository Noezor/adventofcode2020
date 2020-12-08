package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day08.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

type Move struct {
	operation string
	argument  int
}

func parseMove(line string) (move Move) {
	split := strings.Split(line, " ")
	operation := split[0]
	argument, _ := strconv.Atoi(split[1])
	return Move{operation: operation, argument: argument}
}

func findCountAtLoopOrEnd(moves []Move) int {
	isVisited := make(map[int]bool, len(moves))
	indexCurrentMove := 0
	accumulator := 0

	for !isVisited[indexCurrentMove] && indexCurrentMove != len(moves) {
		isVisited[indexCurrentMove] = true
		currentMove := moves[indexCurrentMove]

		if currentMove.operation == "nop" {
			indexCurrentMove += 1
		} else if currentMove.operation == "acc" {
			indexCurrentMove += 1
			accumulator += currentMove.argument
		} else if currentMove.operation == "jmp" {
			indexCurrentMove += currentMove.argument
		}
	}
	return accumulator
}

func isLooping(moves []Move) bool {
	isVisited := make(map[int]bool, len(moves))
	indexCurrentMove := 0
	accumulator := 0

	for !isVisited[indexCurrentMove] && indexCurrentMove != len(moves) {
		isVisited[indexCurrentMove] = true
		currentMove := moves[indexCurrentMove]

		if currentMove.operation == "nop" {
			indexCurrentMove++
		} else if currentMove.operation == "acc" {
			indexCurrentMove++
			accumulator += currentMove.argument
		} else if currentMove.operation == "jmp" {
			indexCurrentMove += currentMove.argument
		}
	}
	return (indexCurrentMove != len(moves))
}

func swapMove(move Move) Move {
	if move.operation == "nop" {
		return Move{operation: "jmp", argument: move.argument}
	} else if move.operation == "jmp" {
		return Move{operation: "nop", argument: move.argument}
	} else {
		return move
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

	moves := make([]Move, len(split))
	for i, s := range split {
		moves[i] = parseMove(s)
	}
	fmt.Println("Part A")
	fmt.Println(findCountAtLoopOrEnd(moves))

	fmt.Println("Part B")
	for i, _ := range moves {
		moves[i] = swapMove(moves[i])
		if !isLooping(moves) {
			fmt.Println(findCountAtLoopOrEnd(moves))
		}
		moves[i] = swapMove(moves[i])
	}
}
