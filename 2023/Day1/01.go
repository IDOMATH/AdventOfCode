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
	var indexes [10]int
	curIdx := strings.Index(str, "zero")
	indexes[0] = curIdx

	curIdx = strings.Index(str, "one")
	indexes[1] = curIdx

	curIdx = strings.Index(str, "two")
	indexes[2] = curIdx

	curIdx = strings.Index(str, "three")
	indexes[3] = curIdx

	curIdx = strings.Index(str, "four")
	indexes[4] = curIdx

	curIdx = strings.Index(str, "five")
	indexes[5] = curIdx

	curIdx = strings.Index(str, "six")
	indexes[6] = curIdx

	curIdx = strings.Index(str, "seven")
	indexes[7] = curIdx

	curIdx = strings.Index(str, "eight")
	indexes[8] = curIdx

	curIdx = strings.Index(str, "nine")
	indexes[9] = curIdx

	min := 1024
	for _, idx := range indexes {
		if idx != -1 && idx < min {
			min = idx
		}
	}
	return min
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
