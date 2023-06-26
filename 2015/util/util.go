package util

import (
	"os"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetFile(path string) (string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(fileBytes), nil
}

func GetLines(file string) []string {
	stringSlice := strings.Split(file, "\n")
	return stringSlice
}

func GetNumberOfFirstCharacters(line string, numberOfCharacters int) []rune {
	firstChars := []rune(line[:numberOfCharacters])

	return firstChars
}
