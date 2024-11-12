package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/idomath/AdventOfCode/util"
)

func main() {
	fmt.Println("hello world")
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
		groupNumber := blankAndGroup[1]

		games := strings.Split(groupAndGames[1], ";")
		for _, game := range games {
			pieceMap := make(map[string]int)
			draws := strings.Split(game, ",")
			for _, draw := range draws {
				colorVal := strings.Split(draw, " ")
				for i := 1; i < len(colorVal); i += 2 {
					numPieces, err := strconv.Atoi(colorVal[i+1])
					if err != nil {
						panic("error getting number of pieces" + err.Error())
					}
					pieceMap[colorVal[i]] = numPieces
				}
			}
		}
		game1 := strings.TrimSpace(games[0])
	}
}
