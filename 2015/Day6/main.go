package main

import (
	"fmt"
	"github.com/idomath/AdventOfCode/2015/util"
	"unicode"
)

type Instruction struct {
	operation string
	xStart    int
	xEnd      int
	yStart    int
	yEnd      int
}

const turnOn = "turn on "
const turnOff = "turn off "
const toggle = "toggle "

type LightGrid [1000][1000]bool
type LightGrid2 [1000][1000]int

func main() {
	RunTest1()
	RunPuzzle1()
	RunTest2()
	RunPuzzle2()
}

func RunTest1() {
	lg := LightGrid{}
	file, err := util.GetFile("./puzzle1_test.txt")
	util.CheckErr(err)
	lines := util.GetLines(file)

	for _, instruction := range lines {
		instr := ParseInstruction([]rune(instruction))
		switch instr.operation {
		case turnOn:
			lg.TurnOn(instr)
		case turnOff:
			lg.TurnOff(instr)
		case toggle:
			lg.Toggle(instr)
		}
	}
	fmt.Println("Lights on: &d", lg.CountLights())
}

func ParseInstruction(instruction []rune) Instruction {
	// instructions have 6 'stages' to them.
	// stage 1 is to set operation i.e. "turn on", "toggle", "turn off"
	// stage 2 is to set x range start
	// stage 3 is to set x range stop
	// stage 4 is to ignore the word "through"
	// stage 5 is to set y range start
	// stage 6 is to set y range stop
	instructionStage := 1
	operation := []rune{}
	xStart := 0
	xEnd := 0
	yStart := 0
	yEnd := 0
	for _, char := range instruction {
		switch instructionStage {
		case 1:
			if !unicode.IsDigit(char) {
				operation = append(operation, char)
			} else {
				xStart = int(char - '0')
				instructionStage++
			}
		case 2:
			if unicode.IsDigit(char) {
				xStart *= 10
				xStart += int(char - '0')
			} else {
				instructionStage++
			}
		case 3:
			if unicode.IsDigit(char) {
				yStart *= 10
				yStart += int(char - '0')
			} else {
				instructionStage++
			}
		case 4:
			if !unicode.IsDigit(char) {
				// Just continue to ignore the characters of "through"
				continue
			} else {
				xEnd = int(char - '0')
				instructionStage++
			}
		case 5:
			if unicode.IsDigit(char) {
				xEnd *= 10
				xEnd += int(char - '0')
			} else {
				instructionStage++
			}
		case 6:
			if unicode.IsDigit(char) {
				yEnd *= 10
				yEnd += int(char - '0')
			} else {
				break
			}
		}
	}

	result := Instruction{
		operation: string(operation),
		xStart:    xStart,
		xEnd:      xEnd,
		yStart:    yStart,
		yEnd:      yEnd,
	}

	return result
}

func (lg *LightGrid) TurnOn(instruction Instruction) {
	for x := instruction.xStart; x <= instruction.xEnd; x++ {
		for y := instruction.yStart; y <= instruction.yEnd; y++ {
			lg[x][y] = true
		}
	}
}
func (lg *LightGrid) TurnOff(instruction Instruction) {
	for x := instruction.xStart; x <= instruction.xEnd; x++ {
		for y := instruction.yStart; y <= instruction.yEnd; y++ {
			lg[x][y] = false
		}
	}
}
func (lg *LightGrid) Toggle(instruction Instruction) {
	for x := instruction.xStart; x <= instruction.xEnd; x++ {
		for y := instruction.yStart; y <= instruction.yEnd; y++ {
			lg[x][y] = !lg[x][y]
		}
	}
}

func (lg *LightGrid) CountLights() int {
	count := 0
	for x := 0; x < len(lg); x++ {
		for y := 0; y < len(lg[x]); y++ {
			if lg[x][y] == true {
				count++
			}
		}
	}
	return count
}

func RunPuzzle1() {
	lg := LightGrid{}
	file, err := util.GetFile("./puzzle1.txt")
	util.CheckErr(err)
	lines := util.GetLines(file)

	for _, instruction := range lines {
		instr := ParseInstruction([]rune(instruction))
		switch instr.operation {
		case turnOn:
			lg.TurnOn(instr)
		case turnOff:
			lg.TurnOff(instr)
		case toggle:
			lg.Toggle(instr)
		}
	}
	fmt.Println("Lights on: &d", lg.CountLights())
}

func (lg *LightGrid2) TurnOn(instruction Instruction) {
	for x := instruction.xStart; x <= instruction.xEnd; x++ {
		for y := instruction.yStart; y <= instruction.yEnd; y++ {
			lg[x][y]++
		}
	}
}
func (lg *LightGrid2) TurnOff(instruction Instruction) {
	for x := instruction.xStart; x <= instruction.xEnd; x++ {
		for y := instruction.yStart; y <= instruction.yEnd; y++ {
			if lg[x][y] != 0 {
				lg[x][y]--
			}
		}
	}
}
func (lg *LightGrid2) Toggle(instruction Instruction) {
	for x := instruction.xStart; x <= instruction.xEnd; x++ {
		for y := instruction.yStart; y <= instruction.yEnd; y++ {
			lg[x][y] += 2
		}
	}
}

func (lg *LightGrid2) CountLights() uint64 {
	var sum uint64
	for x := 0; x < len(lg); x++ {
		for y := 0; y < len(lg[x]); y++ {
			sum += uint64(lg[x][y])
		}
	}
	return sum
}

func RunTest2() {
	lg := LightGrid2{}
	file, err := util.GetFile("./puzzle2_test.txt")
	util.CheckErr(err)
	lines := util.GetLines(file)

	for _, instruction := range lines {
		instr := ParseInstruction([]rune(instruction))
		switch instr.operation {
		case turnOn:
			lg.TurnOn(instr)
		case turnOff:
			lg.TurnOff(instr)
		case toggle:
			lg.Toggle(instr)
		}
	}
	fmt.Println("Lights on: &d", lg.CountLights())
}

func RunPuzzle2() {
	lg := LightGrid2{}
	file, err := util.GetFile("./puzzle1.txt")
	util.CheckErr(err)
	lines := util.GetLines(file)

	for _, instruction := range lines {
		instr := ParseInstruction([]rune(instruction))
		switch instr.operation {
		case turnOn:
			lg.TurnOn(instr)
		case turnOff:
			lg.TurnOff(instr)
		case toggle:
			lg.Toggle(instr)
		}
	}
	fmt.Println("Lights on: &d", lg.CountLights())
}
