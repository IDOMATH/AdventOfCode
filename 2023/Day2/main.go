package main

import (
	"fmt"

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
	for _, line := util.GetLines(file){
		
	}
}
