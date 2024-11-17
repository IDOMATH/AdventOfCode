package main

import (
	"fmt"

	"github.com/idomath/AdventOfCode/util"
)

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
}
