package main

import (
	"fmt"

	"github.com/idomath/AdventOfCode/util"
)

type Location struct {
	Row    int
	Column int
}

type NumberAtLocation struct {
	Num   int
	Start Location
	End   Location
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

	var numberLocations []Location
	for iRow, row := range grid {
		for iColumn, col := range row {
			colNum := int(col - '0')
			if colNum >= 0 && colNum <= 9 {
				numberLocations = append(numberLocations, Location{Row: iRow, Column: iColumn})
			}
		}
	}

	var prevNumLoc Location
	var curNum []int
	for _, numLoc := range numberLocations {

	}

	for _, loc := range symbolLocations {
		if loc.Column-1 >= 0 && loc.Row-1 >= 0 {
			entry := grid[loc.Row-1][loc.Column-1]
			entryNum := int(entry - '0')
			if entryNum >= 0 && entryNum <= 9 {
				//
			}
		}
		if loc.Column > len(grid[0]) {
			continue
		}
		if loc.Row-1 < 0 {
			continue
		}
		if loc.Row < len(grid) {
			continue
		}
	}
}
