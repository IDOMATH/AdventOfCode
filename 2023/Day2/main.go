package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/idomath/AdventOfCode/util"
)

func main() {
	fmt.Println("hello world")
	logic("./test.txt")
}

func logic(filepath string) {
	file, err := util.GetFile(filepath)
	if err != nil {
		panic("error getting file")
	}
	lines := util.GetLines(file)
	var sum = 0
	var redPieces = 12
	var greenPieces = 13
	var bluePieces = 14
	for _, line := range lines {
		groupAndGames := strings.Split(line, ":")
		blankAndGroup := strings.Split(groupAndGames[0], " ")
		groupNumber, err := strconv.Atoi(blankAndGroup[1])
		if err != nil {
			panic("error getting game number")
		}

		legalGame := true
		games := strings.Split(groupAndGames[1], ";")
		for _, game := range games {
			pieceMap := make(map[string]int)
			draws := strings.Split(game, ",")
			for _, draw := range draws {
				colorVal := strings.Split(draw, " ")
				for i := 1; i < len(colorVal); i += 2 {
					numPieces, err := strconv.Atoi(colorVal[i])
					if err != nil {
						panic("error getting number of pieces" + err.Error())
					}
					pieceMap[colorVal[i+1]] = numPieces
				}
				if pieceMap["blue"] > bluePieces ||
					pieceMap["green"] > greenPieces ||
					pieceMap["red"] > redPieces {
					legalGame = false
					break
				}
			}
		}
		if legalGame {
			sum += groupNumber
		}
	}
	fmt.Println(sum)
}
