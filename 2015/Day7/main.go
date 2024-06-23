package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/idomath/AdventOfCode/2015/util"
)

const maxUint16 uint16 = 65535

var outputs map[string]Wire

type Wire struct {
	Value uint16
	IsSet bool
}

func main() {
	file, err := util.GetFile("./puzzle1_test.txt")
	util.CheckErr(err)
	lines := util.GetLines(file)

	outputs = make(map[string]Wire, len(lines))

	// input, output := SplitInputAndOutput(lines[0])
	// fmt.Println(ParseInput(input), output)

	for _, line := range lines {
		// fmt.Println("line: ", line)
		input, output := SplitInputAndOutput(line)
		// fmt.Println("input: ", ParseInput(input), " output: ", output)
		var wire Wire
		wire.Value = ParseInput(input)
		outputs[output] = wire
	}

	for key, output := range outputs {
		fmt.Println(key, ": ", output)
	}
}

func SplitInputAndOutput(instruction string) (input, output string) {
	puts := strings.Split(instruction, "->")
	return puts[0], puts[1]
}

func ParseInput(input string) (value uint16) {
	terms := strings.Split(input, " ")
	if len(terms) == 2 {
		return StrToInt(terms[0])
	}
	if len(terms) == 3 {
		return NotOperator(StrToInt(terms[1]))
	}
	if len(terms) == 4 {
		switch terms[1] {
		case "AND":
			return AndOperator(terms[0], terms[2])
		case "OR":
			return OrOperator(terms[0], terms[2])
		case "LSHIFT":
			return outputs[terms[0]].Value << StrToInt(terms[2])
		case "RSHIFT":
			return outputs[terms[0]].Value >> StrToInt(terms[2])
		}
	}

	return 0
}

func StrToInt(input string) uint16 {
	num, err := strconv.Atoi(input)
	if err != nil {
		return outputs[input].Value
	}
	return uint16(num)
}

func AndOperator(x, y string) uint16 {
	var xi, yi uint16
	i, err := strconv.Atoi(x)
	if err != nil {
		xi = outputs[x].Value
	} else {
		xi = uint16(i)
	}

	i, err = strconv.Atoi(y)
	if err != nil {
		yi = outputs[y].Value
	} else {
		yi = uint16(i)
	}

	return xi & yi
}

func OrOperator(x, y string) uint16 {
	var xi, yi uint16
	i, err := strconv.Atoi(x)
	if err != nil {
		xi = outputs[x].Value
	} else {
		xi = uint16(i)
	}

	i, err = strconv.Atoi(y)
	if err != nil {
		yi = outputs[y].Value
	} else {
		yi = uint16(i)
	}

	return xi ^ yi
}

func NotOperator(num uint16) uint16 {
	return num ^ maxUint16
}
