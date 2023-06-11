package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	RunTestCases()
	fmt.Println(RunPuzzleCase())

	RunPuzzleTwoTestCases()
	fmt.Println(RunPartTwoPuzzle())
}

func GetFloor(directions string) (int, error) {
	floor := 0
	for _, instruction := range directions {
		if string(instruction) == "(" {
			floor++
		} else if string(instruction) == ")" {
			floor--
		} else {
			return 0, errors.New("instruction was not '(' or ')'")
		}
	}
	return floor, nil
}

func RunTestCases() {
	resultFile, err := os.ReadFile("./test_answers.text")
	check(err)
	resultsStringSlice := strings.Split(string(resultFile), "\n")
	results := [9]int{}
	for index, result := range resultsStringSlice {
		results[index], err = strconv.Atoi(result)
		check(err)
	}

	for i, expectedResult := range results {
		input, err := os.ReadFile(fmt.Sprintf("./test%d.text", i+1))
		check(err)
		actualResult, err := GetFloor(string(input))
		check(err)
		if actualResult == expectedResult {
			fmt.Println(fmt.Sprintf("test%d passed", i+1))
		} else {
			fmt.Println(fmt.Sprintf("test%d failed.  expected: %d but got %d", i+1, expectedResult, actualResult))
		}
	}
}

func RunPuzzleCase() int {
	input, err := os.ReadFile("./puzzle_input.text")
	result, err := GetFloor(string(input))
	check(err)
	return result
}

func GetFirstNegativeFloor(directions string) (int, error) {
	floor := 0
	for index, instruction := range directions {
		if string(instruction) == "(" {
			floor++
		} else if string(instruction) == ")" {
			floor--
		} else {
			return 0, errors.New("instruction was not '(' or ')'")
		}
		if floor < 0 {
			return index + 1, nil
		}
	}
	return 0, errors.New("never went into the basement")
}

func RunPuzzleTwoTestCases() {
	resultFile, err := os.ReadFile("./test_answers2.text")
	check(err)
	resultsStringSlice := strings.Split(string(resultFile), "\n")
	results := [2]int{}
	for index, result := range resultsStringSlice {
		results[index], err = strconv.Atoi(result)
		check(err)
	}

	for i, expectedResult := range results {
		input, err := os.ReadFile(fmt.Sprintf("./test%d_puzzle2.text", i+1))
		check(err)
		actualResult, err := GetFirstNegativeFloor(string(input))
		check(err)
		if actualResult == expectedResult {
			fmt.Println(fmt.Sprintf("test%d passed", i+1))
		} else {
			fmt.Println(fmt.Sprintf("test%d failed.  expected: %d but got %d", i+1, expectedResult, actualResult))
		}
	}
}

func RunPartTwoPuzzle() int {
	input, err := os.ReadFile("./puzzle2_input.text")
	result, err := GetFirstNegativeFloor(string(input))
	check(err)
	return result
}
