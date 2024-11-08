package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/idomath/AdventOfCode/util"
)

type IndexNumber struct {
	Index  int
	Number int
}

func main() {
	// runLogic("./test.txt")
	// runLogic("./input01.txt")
	run2Logic("./test2.txt")
	run2Logic("./input01.txt")
}

func transcribeFirstNumber(str string) (int, int) {
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
	val := -1
	for i, idx := range indexes {
		if idx != -1 && idx < min {
			min = idx
			val = i
		}
	}
	return min, val
}

func transcribeLastNumber(str string) (int, int) {
	var indexes [10]int
	curIdx := strings.LastIndex(str, "zero")
	indexes[0] = curIdx

	curIdx = strings.LastIndex(str, "one")
	indexes[1] = curIdx

	curIdx = strings.LastIndex(str, "two")
	indexes[2] = curIdx

	curIdx = strings.LastIndex(str, "three")
	indexes[3] = curIdx

	curIdx = strings.LastIndex(str, "four")
	indexes[4] = curIdx

	curIdx = strings.LastIndex(str, "five")
	indexes[5] = curIdx

	curIdx = strings.LastIndex(str, "six")
	indexes[6] = curIdx

	curIdx = strings.LastIndex(str, "seven")
	indexes[7] = curIdx

	curIdx = strings.LastIndex(str, "eight")
	indexes[8] = curIdx

	curIdx = strings.LastIndex(str, "nine")
	indexes[9] = curIdx

	max := 0
	val := -1
	for i, idx := range indexes {
		if idx != -1 && idx > max {
			max = idx
			val = i
		}
	}
	return max, val
}

func run2Logic(filepath string) {
	file, err := util.GetFile(filepath)
	if err != nil {
		panic("error getting file")
	}
	lines := util.GetLines(file)

	sum := 0
	for lineNum, line := range lines {
		var numbersInLine []IndexNumber
		i, x := transcribeFirstNumber(line)
		numbersInLine = append(numbersInLine, IndexNumber{Index: i, Number: x})
		i, x = transcribeLastNumber(line)
		numbersInLine = append(numbersInLine, IndexNumber{Index: i, Number: x})
		for idx, char := range line {
			num := int(char - '0')
			if num > 9 || num < 0 {
				continue
			}
			numbersInLine = append(numbersInLine, IndexNumber{Index: idx, Number: num})
		}
		lowestIndex := 1024
		leftNumber := -1
		highestIndex := 0
		rightNumber := -1
		for _, nums := range numbersInLine {
			if nums.Index < lowestIndex && nums.Number != -1 {
				lowestIndex = nums.Index
				leftNumber = nums.Number
			}
			if nums.Index > highestIndex && nums.Number != -1 {
				highestIndex = nums.Index
				rightNumber = nums.Number
			}
		}
		sum += (leftNumber * 10) + rightNumber
		fmt.Println(strconv.Itoa(lineNum+1) + " lineVal: " + strconv.Itoa((leftNumber*10)+rightNumber) + ": " + strconv.Itoa(sum))
	}
	fmt.Println(sum)
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
