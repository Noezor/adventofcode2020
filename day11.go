package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day11.input", "Relative file path")

func partA(seats []string) int {
	currentSeats := make([]string, len(seats))
	copy(currentSeats, seats)

	nextSeats := getNextSeatsA(currentSeats)

	for !Equal(nextSeats, currentSeats) {
		currentSeats = nextSeats
		nextSeats = getNextSeatsA(currentSeats)
	}
	return nbOccupiedSeats(currentSeats)
}

func partB(seats []string) int {
	currentSeats := make([]string, len(seats))
	copy(currentSeats, seats)

	nextSeats := getNextSeatsA(currentSeats)

	for !Equal(nextSeats, currentSeats) {
		currentSeats = nextSeats
		nextSeats = getNextSeatsB(currentSeats)
	}
	return nbOccupiedSeats(currentSeats)
}

func nbOccupiedSeats(seats []string) int {
	countOccupied := 0
	for _, rowSeat := range seats {
		for i := 0; i < len(rowSeat); i++ {
			if string(rowSeat[i]) == "#" {
				countOccupied = countOccupied + 1
			}
		}
	}
	return countOccupied
}

func Equal(array1, array2 []string) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}

func replaceAtI(s string, x string, i int) string {
	return s[:i] + x + s[(i+1):]
}

func getNextSeatsA(currentSeats []string) []string {
	nextSeats := make([]string, len(currentSeats))
	copy(nextSeats, currentSeats)
	length := len(currentSeats)
	for i := 0; i < length; i++ {
		width := len(currentSeats[i])
		for j := 0; j < width; j++ {
			nextSeats[i] = replaceAtI(nextSeats[i], getNextValueA(currentSeats, i, j), j)
		}
	}
	return nextSeats
}

func getNextSeatsB(currentSeats []string) []string {
	nextSeats := make([]string, len(currentSeats))
	copy(nextSeats, currentSeats)
	length := len(currentSeats)
	for i := 0; i < length; i++ {
		width := len(currentSeats[i])
		for j := 0; j < width; j++ {
			nextSeats[i] = replaceAtI(nextSeats[i], getNextValueB(currentSeats, i, j), j)
		}
	}
	return nextSeats
}

func getNeighbors(seats []string, i int, j int) []string {
	length := len(seats)
	width := len(seats[i])

	neighbors := make([]string, 0)

	if i-1 >= 0 {
		neighbors = append(neighbors, string(seats[i-1][j]))
		if j-1 >= 0 {
			neighbors = append(neighbors, string(seats[i-1][j-1]))
		}
		if j+1 < width {
			neighbors = append(neighbors, string(seats[i-1][j+1]))
		}
	}
	if i+1 < length {
		neighbors = append(neighbors, string(seats[i+1][j]))
		if j-1 >= 0 {
			neighbors = append(neighbors, string(seats[i+1][j-1]))
		}
		if j+1 < width {
			neighbors = append(neighbors, string(seats[i+1][j+1]))
		}
	}
	if j-1 >= 0 {
		neighbors = append(neighbors, string(seats[i][j-1]))

	}
	if j+1 < width {
		neighbors = append(neighbors, string(seats[i][j+1]))
	}
	return neighbors
}

func getVisible(seats []string, i int, j int) []string {
	visible := make([]string, 0)

	for verticalStep := -1; verticalStep <= +1; verticalStep++ {
		for horizontalStep := -1; horizontalStep <= +1; horizontalStep++ {
			if verticalStep != 0 || horizontalStep != 0 {
				distance := 1
				seenVisible := false
				for exists(seats, i+verticalStep*distance, j+horizontalStep*distance) && !seenVisible {
					if seen := string(seats[i+verticalStep*distance][j+horizontalStep*distance]); seen == "#" || seen == "L" {
						seenVisible = true
						visible = append(visible, seen)
					}
					distance = distance + 1
				}
			}
		}
	}
	return visible
}

func exists(seats []string, i int, j int) bool {
	length := len(seats)
	if !(0 <= i && i < length) {
		return false
	}
	width := len(seats[i])
	return 0 <= j && j < width
}

func count(s []string) map[string]int {
	countValues := make(map[string]int)
	for _, v := range s {
		countValues[v] = countValues[v] + 1
	}
	return countValues
}

func getNextValueA(seats []string, i int, j int) string {
	comparedValue := string(seats[i][j])
	if comparedValue == "." {
		return comparedValue
	}
	countNeighbors := count(getNeighbors(seats, i, j))
	if comparedValue == "L" && countNeighbors["#"] == 0 {
		return "#"
	}
	if comparedValue == "#" && countNeighbors["#"] >= 4 {
		return "L"
	}
	return comparedValue
}

func getNextValueB(seats []string, i int, j int) string {
	comparedValue := string(seats[i][j])
	if comparedValue == "." {
		return comparedValue
	}
	countNeighbors := count(getVisible(seats, i, j))
	if comparedValue == "L" && countNeighbors["#"] == 0 {
		return "#"
	}
	if comparedValue == "#" && countNeighbors["#"] >= 5 {
		return "L"
	}
	return comparedValue
}

func parseString(contents string) []string {
	split := strings.Split(contents, "\n")

	seats := make([]string, len(split))
	for i, s := range split {
		seats[i] = s
	}
	return seats
}

var TestString = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
var expectedOutputA = 37
var expectedOutputB = 26

func main() {
	flag.Parse()
	contents := TestString

	if result := partA(parseString(contents)); result != expectedOutputA {
		fmt.Println("Error, expected", expectedOutputA, "Got ", result)
		return
	}
	if result := partB(parseString(contents)); result != expectedOutputB {
		fmt.Println("Error, expected", expectedOutputB, "Got ", result)
		return
	}
	fmt.Println("Passed tests")

	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents = string(bytes)
	fmt.Println(partA(parseString(contents)))
	fmt.Println(partB(parseString(contents)))
}
