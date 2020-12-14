package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day14.input", "Relative file path")

type instruction struct {
	action string
	value  string
}

func partA(instructions []instruction) int {
	mask := make(map[int]int)

	mem := make(map[int]int)
	for _, ins := range instructions {
		if ins.action == "mask" {
			mask = getBinValue(ins)
		} else {
			changeMem(mem, ins, mask)
		}
	}
	sumMemory := 0
	for _, value := range mem {
		sumMemory += value
	}
	return sumMemory
}

func partB(instructions []instruction) int {
	mask := make(map[int]int)

	mem := make(map[int]int)
	for _, ins := range instructions {
		if ins.action == "mask" {
			mask = getBinValue(ins)
		} else {
			value, err := strconv.Atoi(ins.value)
			if err != nil {
				fmt.Println("Could not parse ", ins.action)
				return -1
			}

			addresses := getAddresses(ins, mask)
			for _, address := range addresses {
				mem[address] = value
			}
		}
	}
	sumMemory := 0
	for _, value := range mem {
		sumMemory += value
	}
	return sumMemory
}

func changeMem(mem map[int]int, ins instruction, mask map[int]int) {
	memAdress, err := strconv.Atoi(ins.action[4:(len(ins.action) - 1)])
	if err != nil {
		fmt.Println("Could not parse ", ins.action)
		return
	}
	newMemValue, err := strconv.Atoi(ins.value)
	if err != nil {
		fmt.Println("Could not parse ", ins.value)
		return
	}
	mem[memAdress] = applyMask(newMemValue, mask)
}

func getAddresses(ins instruction, mask map[int]int) []int {
	memAdress, err := strconv.Atoi(ins.action[4:(len(ins.action) - 1)])
	if err != nil {
		fmt.Println("Could not parse ", ins.action)
		return []int{}
	}
	return getAllValues(memAdress, mask)
}

func getAllValues(value int, mask map[int]int) []int {
	overwrittenValue := overwriteValue(value, mask)

	possibleValues := []int{overwrittenValue}
	for currentBinaryIndex := 0; currentBinaryIndex <= 35; currentBinaryIndex++ {
		decimalValue := 1 << currentBinaryIndex
		if _, ok := mask[currentBinaryIndex]; !ok {
			lenValuesBeforeDup := len(possibleValues)
			for i := 0; i < lenValuesBeforeDup; i++ {
				newValue := possibleValues[i] + decimalValue
				possibleValues = append(possibleValues, newValue)
			}
		}
	}
	return possibleValues
}

func overwriteValue(value int, mask map[int]int) int {
	newValue := 0
	for currentBinaryIndex := 0; currentBinaryIndex <= 35; currentBinaryIndex++ {
		decimalValue := 1 << currentBinaryIndex
		if val, ok := mask[currentBinaryIndex]; ok {
			if val == 1 {
				newValue += decimalValue
			} else if val == 0 {
				newValue += (decimalValue & value)
			}
		}
	}
	return newValue
}

func applyMask(value int, mask map[int]int) int {
	newValue := 0
	for currentBinaryIndex := 0; currentBinaryIndex <= 35; currentBinaryIndex++ {
		decimalValue := 1 << currentBinaryIndex
		if val, ok := mask[currentBinaryIndex]; ok {
			newValue += decimalValue * val
		} else {
			newValue += (value & decimalValue)
		}
	}
	return newValue
}

func getBinValue(ins instruction) map[int]int {
	if ins.action != "mask" {
		fmt.Println("Wrong input type", ins.action)
		return make(map[int]int, 0)
	}

	binaryMask := make(map[int]int, 36)
	for i := 0; i < len(ins.value); i++ {
		if string(ins.value[35-i]) != "X" {
			value, err := strconv.Atoi(string(ins.value[35-i]))
			if err != nil {
				fmt.Println("Failed Parsing of", ins.value)
			}
			binaryMask[i] = value
		}
	}
	return binaryMask
}

func toBinary(ins instruction) map[int]int {
	return make(map[int]int)
}

func parseString(contents string) []instruction {
	split := strings.Split(contents, "\n")
	instructions := make([]instruction, len(split))

	for i, line := range split {
		splitLine := strings.Split(line, " ")
		action, value := splitLine[0], splitLine[2]
		instructions[i] = instruction{action: action, value: value}
	}
	return instructions
}

var testString = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
var expectedOutput = 165

var testStringB = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
var expectedOutputB = 208

func main() {
	flag.Parse()
	contents := testString

	if result := partA(parseString(contents)); result != expectedOutput {
		fmt.Println("Error, expected", expectedOutput, "Got ", result)
		return
	}
	contents = testStringB
	if result := partB(parseString(contents)); result != expectedOutputB {
		fmt.Println("Error, expected", expectedOutputB, "Got ", result)
		return
	}
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
