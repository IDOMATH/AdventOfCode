package main

import (
	"fmt"
	"strings"

	"github.com/idomath/AdventOfCode/util"
)

func main() {
	runLogic("./test.txt")
	runLogic("./input01.txt")
}

func transcribeFirstNumber(str string) int {
	var zeroIdx, oneIdx, twoIdx, threeIdx, fourIdx, fiveIdx, sixIdx, sevenIdx, eightIdx, nineIdx, curIdx int
	curIdx = strings.Index(str, "zero")
	if curIdx != -1 {
		zeroIdx = curIdx
	}
	curIdx = strings.Index(str, "one")
	if curIdx != -1 {
		oneIdx = curIdx
	}
	curIdx = strings.Index(str, "two")
	if curIdx != -1 {
		twoIdx = curIdx
	}
	curIdx = strings.Index(str, "three")
	if curIdx != -1 {
		threeIdx = curIdx
	}
	curIdx = strings.Index(str, "four")
	if curIdx != -1 {
		fourIdx = curIdx
	}
	curIdx = strings.Index(str, "five")
	if curIdx != -1 {
		fiveIdx = curIdx
	}
	curIdx = strings.Index(str, "six")
	if curIdx != -1 {
		sixIdx = curIdx
	}
	curIdx = strings.Index(str, "seven")
	if curIdx != -1 {
		sevenIdx = curIdx
	}
	curIdx = strings.Index(str, "eight")
	if curIdx != -1 {
		eightIdx = curIdx
	}
	curIdx = strings.Index(str, "nine")
	if curIdx != -1 {
		nineIdx = curIdx
	}
}

func runLogic(filepath string) {
	file, err := util.GetFile(filepath)
	if err != nil {
		panic("error getting file")
	}
	lines := util.GetLines(file)

	var numbersInLine [][]int
	for _, line := range lines {
		var numbers []int
		for _, char := range line {
			num := int(char - '0')
			if num > 9 || num < 0 {
				continue
			}
			numbers = append(numbers, num)
		}
		numbersInLine = append(numbersInLine, numbers)
	}

	var sum = 0
	for _, ns := range numbersInLine {
		sum += (ns[0] * 10) + ns[len(ns)-1]
	}
	fmt.Println(sum)
}
