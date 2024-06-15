package main

import (
	"fmt"
	"strconv"
	"strings"
)

const maxUint16 uint16 = 65535

func main() {
	// file, err := util.GetFile("./puzzle1_test.txt")
	// util.CheckErr(err)
	// lines := util.GetLines(file)

	// outputs := make(map[string]uint16, len(lines))

	fmt.Println(NotOperator(1))

	// for _, line := range lines {
	// 	fmt.Println("line: ", line)
	// 	input, output := SplitInputAndOutput(line)
	// 	fmt.Println("input: ", input, " output: ", output)
	// }
}

func SplitInputAndOutput(instruction string) (input, output string) {
	puts := strings.Split(instruction, "->")
	return puts[0], puts[1]
}

func ParseInput(input string, outs map[string]uint16) (value uint16) {
	terms := strings.Split(input, " ")
	if len(terms) == 1 {
		return StrToInt(terms[0])
	}
	if len(terms) == 2 {
		return NotOperator(StrToInt(terms[1]))
	}

	return 0
}

func StrToInt(input string) uint16 {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic("could not get input number")
	}
	return uint16(num)
}

func NotOperator(num uint16) uint16 {
	return num ^ maxUint16
}
