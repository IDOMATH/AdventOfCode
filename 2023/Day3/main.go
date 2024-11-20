package main

import (
	"fmt"

	"github.com/idomath/AdventOfCode/util"
)

type Location struct {
	Row    int
	Column int
}

func main() {
	fmt.Println("Hello world")
}

func logic(filepath string) {
	file, err := util.GetFile(filepath)
	if err != nil {
		panic("error getting file")
	}

	lines := util.GetLines(file)

	var grid [][]rune

	for _, line := range lines {
		var gridRow []rune
		for _, char := range line {
			gridRow = append(gridRow, char)

		}
		grid = append(grid, gridRow)
	}

	// Get the locations of the symbols
	var symbolLocations []Location
	for iRow, row := range grid {
		for iCol, col := range row {
			colNum := int(col - '0')
			if colNum >= 0 && colNum <= 9 {
				continue
			}
			if col != '.' {
				symbolLocations = append(symbolLocations, Location{Row: iRow, Column: iCol})
			}
		}
	}
}
