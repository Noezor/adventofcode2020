package main

// template by LFJ

import(
	"flag",
	"fmt",
	"io/ioutil",
	"strconv",
	"strings"
)


var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path")
var partB = flag.String("partB", false, "Using part B logic")

func main(){
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	sum := 0
	split := strings.Split(contents, "\n")
	for _,s := range split {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
	}
	fmt.Println(sum)
}