package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day04.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

var manadatoryKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func verifyPassport(passport map[string]string, mandatoryKeys []string) bool {
	presenceMandatoryKeys := presenceKeys(mandatoryKeys)
	for key, value := range passport {
		if _, ok := presenceMandatoryKeys[key]; ok {
			presenceMandatoryKeys[key] = verifyKey(key, value)
		}
	}
	return allTrue(presenceMandatoryKeys)
}

func presenceKeys(keys []string) map[string]bool {
	presenceKeys := make(map[string]bool, len(keys))
	for _, key := range keys {
		presenceKeys[key] = false
	}
	return presenceKeys
}

func allTrue(booleanMap map[string]bool) bool {
	for key := range booleanMap {
		if !booleanMap[key] {
			return false
		}
	}
	return true
}

func splitLinesPassport(concatStrings string) map[string]string {
	dictionnary := make(map[string]string, 8)
	split := strings.Split(concatStrings, "\n")
	for _, s := range split {
		dictionnary = addLineToDictionnary(s, dictionnary)
	}
	return dictionnary
}

func addLineToDictionnary(line string, dictionnary map[string]string) map[string]string {
	split := strings.Split(line, " ")
	for _, keyValue := range split {
		dictionnary = addKeyValueToDictionnary(keyValue, dictionnary)
	}
	return dictionnary
}

func verifyKey(key string, value string) bool {
	if key == "byr" {
		return "1920" <= value && value <= "2002" && len(value) == 4
	}
	if key == "iyr" {
		return "2010" <= value && value <= "2020" && len(value) == 4
	}
	if key == "eyr" {
		return "2020" <= value && value <= "2030" && len(value) == 4
	}
	if key == "hgt" {
		return validateHeight(value)
	}
	if key == "hcl" {
		return validateHCL(value)
	}
	if key == "ecl" {
		return validateECL(value)
	}
	if key == "pid" {
		if "000000000" <= value && value <= "999999999" && len(value) == 9 {
			return true
		}

	}
	return false
}

func validateHeight(value string) bool {
	if strings.HasSuffix(value, "in") {
		size := value[:2]
		return "59" <= size && size <= "76" && len(value) == 4
	}
	if strings.HasSuffix(value, "cm") {
		size := value[:3]
		return "150" <= size && size <= "193" && len(value) == 5
	}
	return false
}

func validateECL(value string) bool {
	possibleValues := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, e := range possibleValues {
		if e == value {
			return true
		}
	}
	return false
}

func validateHCL(value string) bool {
	if string(value[0]) != "#" {
		return false
	}
	for _, char := range value[1:] {
		if !isValidCharacterECL(string(char)) {
			return false
		}
	}
	return len(value) == 7
}

func isValidCharacterECL(char string) bool {
	return ("0" <= char && char <= "9") || ("a" <= char && char <= "f")
}

func addKeyValueToDictionnary(keyValue string, dictionnary map[string]string) map[string]string {
	split := strings.Split(keyValue, ":")
	if len(split) == 2 {
		key := split[0]
		value := split[1]
		dictionnary[key] = value
	}
	return dictionnary
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	validPassports := 0
	split := strings.Split(contents, "\n")
	currentPassportString := ""
	for _, s := range split {
		if s == "" {
			if verifyPassport(splitLinesPassport(currentPassportString), manadatoryKeys) {
				validPassports++
			}
			currentPassportString = ""
		} else {
			currentPassportString = currentPassportString + "\n" + s
		}
	}
	if verifyPassport(splitLinesPassport(currentPassportString), manadatoryKeys) {
		validPassports++
	}
	fmt.Println(validPassports)
}
