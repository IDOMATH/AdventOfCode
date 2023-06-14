package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type CurrentPosition struct {
	x int
	y int
}

func main() {
	RunTestCases()
	fmt.Println(RunPuzzleCase())
	fmt.Println(RunPuzzle2TestCaseX(1))
	fmt.Println(RunPuzzle2TestCaseX(2))
	fmt.Println(RunPuzzle2TestCaseX(3))
	fmt.Println(RunPuzzle2())
}

func NewCurrentPosition() *CurrentPosition {
	return &CurrentPosition{
		x: 0,
		y: 0,
	}
}

func (c *CurrentPosition) MovePosition(deltaX, deltaY int) {
	c.y += deltaY
	c.x += deltaX
}

func GetDelta(direction string) (deltaX, deltaY int, err error) {
	switch direction {
	case "^":
		return 0, 1, nil
	case ">":
		return 1, 0, nil
	case "v":
		return 0, -1, nil
	case "<":
		return -1, 0, nil
	default:
		return 0, 0, errors.New("input must be one of '^', '>', 'v', '<'")
	}
}

func DeliverPresent(curPos CurrentPosition, deliveredTo map[string]int) {
	key := fmt.Sprintf("%d,%d", curPos.x, curPos.y)
	deliveredTo[key] = 1
}

func GetNumberOfHousesDeliveredTo(deliveredTo map[string]int) int {
	return len(deliveredTo)
}

func RunTestCases() {
	resultFile, err := os.ReadFile("./puzzle1_answers.text")
	check(err)
	resultsStringSlice := strings.Split(string(resultFile), "\n")
	results := [3]int{}
	for index, result := range resultsStringSlice {
		results[index], err = strconv.Atoi(result)
		check(err)
	}

	for i, expectedResult := range results {
		input, err := os.ReadFile(fmt.Sprintf("./puzzle1_test%d.text", i+1))
		check(err)

		currentPosition := NewCurrentPosition()
		deliveryMap := make(map[string]int)

		DeliverPresent(*currentPosition, deliveryMap)

		for _, direction := range string(input) {
			deltaX, deltaY, err := GetDelta(string(direction))
			check(err)
			currentPosition.MovePosition(deltaX, deltaY)
			DeliverPresent(*currentPosition, deliveryMap)
		}

		actualResult := GetNumberOfHousesDeliveredTo(deliveryMap)
		if actualResult == expectedResult {
			fmt.Println(fmt.Sprintf("test%d passed", i+1))
		} else {
			fmt.Println(fmt.Sprintf("test%d failed.  expected: %d but got %d", i+1, expectedResult, actualResult))
		}
	}
}

func RunPuzzleCase() int {
	input, err := os.ReadFile("./puzzle1.text")
	check(err)

	currentPosition := NewCurrentPosition()
	deliveryMap := make(map[string]int)

	DeliverPresent(*currentPosition, deliveryMap)

	for _, direction := range string(input) {
		deltaX, deltaY, err := GetDelta(string(direction))
		check(err)
		currentPosition.MovePosition(deltaX, deltaY)
		DeliverPresent(*currentPosition, deliveryMap)
	}

	return GetNumberOfHousesDeliveredTo(deliveryMap)
}

func GetAnswerForPuzzleXTestY(puzzleNumber, testNumber int) int {
	resultFile, err := os.ReadFile(fmt.Sprintf("./puzzle%d_answers.text", puzzleNumber))
	check(err)
	resultsStringSlice := strings.Split(string(resultFile), "\n")
	answer, err := strconv.Atoi(resultsStringSlice[testNumber-1])
	check(err)
	return answer
}

func RunPuzzle2TestCaseX(testNumber int) bool {
	input, err := os.ReadFile(fmt.Sprintf("./puzzle2_test%d.text", testNumber))
	check(err)

	currentRealPosition := NewCurrentPosition()
	currentRoboPosition := NewCurrentPosition()

	deliveryMap := make(map[string]int)

	DeliverPresent(*currentRealPosition, deliveryMap)
	DeliverPresent(*currentRoboPosition, deliveryMap)

	for index, direction := range string(input) {
		deltaX, deltaY, err := GetDelta(string(direction))
		check(err)
		if index%2 == 0 {
			currentRealPosition.MovePosition(deltaX, deltaY)
			DeliverPresent(*currentRealPosition, deliveryMap)
		} else {
			currentRoboPosition.MovePosition(deltaX, deltaY)
			DeliverPresent(*currentRoboPosition, deliveryMap)
		}
	}
	return GetNumberOfHousesDeliveredTo(deliveryMap) == GetAnswerForPuzzleXTestY(2, testNumber)
}

func RunPuzzle2() int {
	input, err := os.ReadFile("./puzzle2.text")
	check(err)

	currentRealPosition := NewCurrentPosition()
	currentRoboPosition := NewCurrentPosition()

	deliveryMap := make(map[string]int)

	DeliverPresent(*currentRealPosition, deliveryMap)
	DeliverPresent(*currentRoboPosition, deliveryMap)

	for index, direction := range string(input) {
		deltaX, deltaY, err := GetDelta(string(direction))
		check(err)
		if index%2 == 0 {
			currentRealPosition.MovePosition(deltaX, deltaY)
			DeliverPresent(*currentRealPosition, deliveryMap)
		} else {
			currentRoboPosition.MovePosition(deltaX, deltaY)
			DeliverPresent(*currentRoboPosition, deliveryMap)
		}
	}
	return GetNumberOfHousesDeliveredTo(deliveryMap)
}
