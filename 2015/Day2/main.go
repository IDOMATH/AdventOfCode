package main

import (
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
	RunPuzzleOneTestCases()
	RunPuzzleOneTestSum()
	fmt.Println(RunPuzzleOne())
	RunPuzzleTwoTestSum()
	fmt.Println(RunPuzzleTwo())
}

func FindRequiredWrappingPaper(length, width, height int) int {
	return FindSurfaceArea(length, width, height) + FindSmallestSide(length, width, height)
}

func FindSurfaceArea(length, width, height int) int {
	return 2 * ((length * width) + (length * height) + (height * width))
}

func FindSmallestSide(length, width, height int) int {
	face1 := length * width
	face2 := length * height
	face3 := width * height

	smallestFace := face1
	if face2 < smallestFace {
		smallestFace = face2
	}
	if face3 < smallestFace {
		smallestFace = face3
	}

	return smallestFace
}

func FindRibbonRequired(length, width, height int) int {
	perimeter1 := FindPerimeter(length, width)
	perimeter2 := FindPerimeter(length, height)
	perimeter3 := FindPerimeter(width, height)

	smallestPerimeter := perimeter1
	if perimeter2 < smallestPerimeter {
		smallestPerimeter = perimeter2
	}
	if perimeter3 < smallestPerimeter {
		smallestPerimeter = perimeter3
	}

	return smallestPerimeter + FindCubicVolume(length, width, height)
}

func FindCubicVolume(length, width, height int) int {
	return length * width * height
}

func FindPerimeter(length, width int) int {
	return 2 * (length + width)
}

func RunPuzzleOneTestCases() {
	resultFile, err := os.ReadFile("./puzzle1_answers.text")
	check(err)
	resultsStringSlice := strings.Split(string(resultFile), "\n")
	results := [2]int{}
	for index, result := range resultsStringSlice {
		results[index], err = strconv.Atoi(result)
		check(err)
	}

	for i, expectedResult := range results {
		inputFile, err := os.ReadFile(fmt.Sprintf("./puzzle1_test%d.text", i+1))
		check(err)
		inputStringSlice := strings.Split(string(inputFile), "\n")
		inputIntSlice := [3]int{}
		for index, value := range inputStringSlice {
			inputIntSlice[index], err = strconv.Atoi(value)
			check(err)
		}
		actualResult := FindRequiredWrappingPaper(inputIntSlice[0], inputIntSlice[1], inputIntSlice[2])

		if actualResult == expectedResult {
			fmt.Println(fmt.Sprintf("test%d passed", i+1))
		} else {
			fmt.Println(fmt.Sprintf("test%d failed.  expected: %d but got %d", i+1, expectedResult, actualResult))
		}
	}
}

func RunPuzzleOneTestSum() {
	resultFile, err := os.ReadFile("./puzzle1_answer.text")
	check(err)
	result, err := strconv.Atoi(string(resultFile))
	check(err)

	inputFile, err := os.ReadFile("./puzzle1_tests.txt.text")
	check(err)
	inputRows := strings.Split(string(inputFile), "\n")
	var inputs [2][3]int
	for index, value := range inputRows {
		lwhStringSlice := strings.Split(value, "x")
		for measurementIdx, measurementValue := range lwhStringSlice {
			inputs[index][measurementIdx], err = strconv.Atoi(measurementValue)
			check(err)
		}
	}

	sum := 0
	for _, box := range inputs {
		sum += FindRequiredWrappingPaper(box[0], box[1], box[2])
	}
	if result == sum {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}
}

func RunPuzzleOne() int {
	inputFile, err := os.ReadFile("./puzzle1_input.text")
	check(err)
	inputRows := strings.Split(string(inputFile), "\n")
	var inputs [1000][3]int
	for index, value := range inputRows {
		lwhStringSlice := strings.Split(value, "x")
		for measurementIdx, measurementValue := range lwhStringSlice {
			inputs[index][measurementIdx], err = strconv.Atoi(measurementValue)
			check(err)
		}
	}

	sum := 0
	for _, box := range inputs {
		sum += FindRequiredWrappingPaper(box[0], box[1], box[2])
	}
	return sum
}

func RunPuzzleTwoTestSum() {
	resultFile, err := os.ReadFile("./puzzle2_answer.text")
	check(err)
	result, err := strconv.Atoi(string(resultFile))
	check(err)

	inputFile, err := os.ReadFile("./puzzle2_tests.text")
	check(err)
	inputRows := strings.Split(string(inputFile), "\n")
	var inputs [2][3]int
	for index, value := range inputRows {
		lwhStringSlice := strings.Split(value, "x")
		for measurementIdx, measurementValue := range lwhStringSlice {
			inputs[index][measurementIdx], err = strconv.Atoi(measurementValue)
			check(err)
		}
	}

	sum := 0
	for _, box := range inputs {
		sum += FindRibbonRequired(box[0], box[1], box[2])
	}
	if result == sum {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}
}

func RunPuzzleTwo() int {
	inputFile, err := os.ReadFile("./puzzle2_input.text")
	check(err)
	inputRows := strings.Split(string(inputFile), "\n")
	var inputs [1000][3]int
	for index, value := range inputRows {
		lwhStringSlice := strings.Split(value, "x")
		for measurementIdx, measurementValue := range lwhStringSlice {
			inputs[index][measurementIdx], err = strconv.Atoi(measurementValue)
			check(err)
		}
	}

	sum := 0
	for _, box := range inputs {
		sum += FindRibbonRequired(box[0], box[1], box[2])
	}
	return sum
}
