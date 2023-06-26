package main

import (
	"fmt"
	"github.com/idomath/AdventOfCode/2015/util"
	"os"
)

func main() {
	RunTests()
	fmt.Println(RunPuzzle())
	RunTestsPart2()
	fmt.Println(RunPuzzle2())
}

func IsVowel(char string) bool {
	const vowels = "aeiou"
	for i := 0; i < len(vowels); i++ {
		if char[0] == vowels[i] {
			return true
		}
	}
	return false
}

func IsDoubleLetter(str string) bool {
	if str[0] == str[1] {
		return true
	}
	return false
}

func IsBadString(str string) bool {
	badStrings := [4]string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, badString := range badStrings {
		if str == badString {
			return true
		}
	}
	return false
}

func CheckString(str string) bool {
	isNice := false
	hasDoubleLetter := false
	vowelCount := 0
	oldChar := " "
	for _, char := range str {
		if IsVowel(string(char)) {
			vowelCount++
		}
		if IsDoubleLetter(fmt.Sprintf(oldChar + string(char))) {
			hasDoubleLetter = true
		}
		if IsBadString(fmt.Sprintf(oldChar + string(char))) {
			return false
		}
		oldChar = string(char)
	}
	if vowelCount >= 3 && hasDoubleLetter == true {
		isNice = true
	}
	return isNice
}

func RunTests() {
	file, err := util.GetFile("./puzzle1_tests.txt")
	util.CheckErr(err)
	testInputs := util.GetLines(file)

	for i, input := range testInputs {
		fmt.Printf("Is test %d nice? %v\n", i+1, CheckString(input))
	}
}

func RunPuzzle() int {
	file, err := util.GetFile("./puzzle1.txt")
	util.CheckErr(err)
	testInputs := util.GetLines(file)
	niceCount := 0

	for _, input := range testInputs {
		if CheckString(input) {
			niceCount++
		}
	}
	return niceCount
}

func HasDoublePair(tuples [][2]rune) bool {
	for i := 0; i < len(tuples); i++ {
		for j := i + 2; j < len(tuples); j++ {
			if (string(tuples[i][0]) == string(tuples[j][0])) && (string(tuples[i][1]) == string(tuples[j][1])) {
				//fmt.Printf("first rune: %s, second rune: %s", string(tuples[i][0]), string(tuples[i][1]))
				//fmt.Printf("Pair is: %v, %v \nat indexes: %d, %d\n", tuples[i][0:1], tuples[j][0:1], i, j)
				return true
			}
		}
	}
	return false
}

func IsRepeatedWithGap(str [3]rune) bool {
	if str[0] == str[2] {
		return true
	}
	return false
}

func CheckStringPart2(str []rune) bool {
	isNice := false
	hasRepeatingWithGap := false
	hasPair := false
	pairs := [][2]rune{}
	for i := 0; i < len(str)-1; i++ {
		pair := [2]rune{str[i], str[i+1]}
		pairs = append(pairs, pair)
		if i < len(str)-2 {
			tripleRune := [3]rune{str[i], str[i+1], str[i+2]}
			if IsRepeatedWithGap(tripleRune) {
				hasRepeatingWithGap = true
			}
		}
	}

	hasPair = HasDoublePair(pairs)

	if hasRepeatingWithGap && hasPair {
		isNice = true
	}
	return isNice
}

func RunTestsPart2() {
	file, err := os.ReadFile("./puzzle2_tests.txt")
	util.CheckErr(err)
	testInputs := util.GetLines(string(file))

	for i, input := range testInputs {
		fmt.Printf("Is test %d nice? %v\n", i+1, CheckStringPart2([]rune(input)))
	}
}

func RunPuzzle2() int {
	file, err := util.GetFile("./puzzle2.txt")
	util.CheckErr(err)
	testInputs := util.GetLines(file)
	niceCount := 0

	for _, input := range testInputs {
		if CheckStringPart2([]rune(input)) {
			niceCount++
		}
	}
	return niceCount
}
