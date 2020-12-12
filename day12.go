package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day12.input", "Relative file path")

type move struct {
	action string
	value  int
}

func partA(moves []move) int {
	return manathanDistanceToOrigin(executeMovesA(moves))
}

func executeMovesA(moves []move) (finalX int, finalY int) {
	x, y := 0, 0
	currentAngle := 90
	directions := [...]string{"N", "E", "S", "W"}
	for _, move := range moves {
		if move.action == "R" {
			currentAngle += move.value
		} else if move.action == "L" {
			currentAngle -= move.value
		} else {
			var directionMove string
			magnitudeMove := move.value
			if move.action == "F" {
				directionMove = directions[mod(currentAngle/90, 4)]
			} else {
				directionMove = move.action
			}
			x, y = nextPositionA(x, y, directionMove, magnitudeMove)
		}
	}
	return x, y
}

func partB(moves []move) int {
	return manathanDistanceToOrigin(executeMovesB(moves))
}

func executeMovesB(moves []move) (finalX int, finalY int) {
	x, y := 0, 0
	polX, polY := 10, 1
	for _, move := range moves {
		if move.action == "F" {
			x, y = nextPositionB(x, y, polX, polY, move.value)
		} else {
			polX, polY = movePol(polX, polY, move.action, move.value)
		}
	}
	return x, y
}

func nextPositionA(x int, y int, directionMove string, magnitudeMove int) (int, int) {
	if directionMove == "N" {
		return x, y + magnitudeMove
	} else if directionMove == "S" {
		return x, y - magnitudeMove
	} else if directionMove == "E" {
		return x + magnitudeMove, y
	} else if directionMove == "W" {
		return x - magnitudeMove, y
	} else {
		fmt.Println("wrong direction", directionMove)
		return 0, 0
	}
}

func nextPositionB(x, y, polX, polY, nbForward int) (int, int) {
	return x + polX*nbForward, y + polY*nbForward
}

func movePol(x int, y int, action string, value int) (int, int) {
	if action == "N" {
		return x, y + value
	} else if action == "S" {
		return x, y - value
	} else if action == "E" {
		return x + value, y
	} else if action == "W" {
		return x - value, y
	} else {
		nbTurn := value / 90
		if action == "R" {
			for i := 0; i < nbTurn; i++ {
				x, y = y, -x
			}
		} else if action == "L" {
			for i := 0; i < nbTurn; i++ {
				x, y = -y, x
			}
		}
		return x, y
	}
	fmt.Println("wrong direction", action)
	return 0, 0
}

func manathanDistanceToOrigin(x int, y int) int {
	return abs(x) + abs(y)
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func mod(a int, b int) int {
	normalMod := a % b
	for normalMod < 0 {
		a = a + b
		normalMod = a % b
	}
	return normalMod
}

func parseString(contents string) []move {
	split := strings.Split(contents, "\n")
	moves := make([]move, len(split))
	for i, s := range split {
		value, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		action := string(s[0])
		moves[i] = move{action: action, value: value}
	}
	return moves
}

var testString = `F10
N3
F7
R90
F11`
var expectedOutput = 25
var expectedOutputB = 286

func main() {
	flag.Parse()
	contents := testString

	if result := partA(parseString(contents)); result != expectedOutput {
		fmt.Println("Error, expected", expectedOutput, "Got ", result)
		return
	}
	if result := partB(parseString(contents)); result != expectedOutputB {
		fmt.Println("Error, expected", expectedOutputB, "Got ", result)
		return
	}
	fmt.Println("passed tests")

	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents = string(bytes)
	fmt.Println(partA(parseString(contents)))
	fmt.Println(partB(parseString(contents)))
}
