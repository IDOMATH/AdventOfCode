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

func transcribeNumber(str string) int {
	var zeroIdx, oneIdx, twoIdx, threeIdx, fourIdx, fiveIdx, sixIdx, sevenIdx, eightIdx, nineIdx, curIdx int
	curIdx = strings.Index(str, "zero")
	if curIdx != -1 {
		zeroIdx = curIdx
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
