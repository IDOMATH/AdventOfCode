package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

// If I was smarter or this took too long, I would look into the actual hashing algorithm
// and reverse engineer a solution, but instead we're just brute forcing for an answer
// because it really doesn't take long to execute.
func main() {
	fmt.Println(CheckFirst5Characters(GetSumFromSecretAndAnswer("abcdef", "609043")))
	fmt.Println(CheckFirst5Characters(GetSumFromSecretAndAnswer("pqrstuv", "1048970")))
	fmt.Println(RunPuzzle1("bgvyzdsv"))
	fmt.Println(RunPuzzle2("bgvyzdsv"))
}

func GetSumFromSecretAndAnswer(secret, answer string) string {
	hash := md5.New()
	io.WriteString(hash, secret)
	io.WriteString(hash, answer)
	return string(hash.Sum(nil))
}

func CheckFirst5Characters(hashString string) bool {
	const fiveZeroes = "00000"
	for i := 0; i < 5; i++ {
		hash := fmt.Sprintf("%x", hashString)
		if hash[i] != fiveZeroes[i] {
			return false
		}

	}
	return true
}

func RunPuzzle1(secret string) int {
	i := 0
	hash := GetSumFromSecretAndAnswer(secret, strconv.Itoa(i))
	for failed := true; failed; failed = !CheckForLeadingZeroes(hash, 5) {
		i++
		hash = GetSumFromSecretAndAnswer(secret, strconv.Itoa(i))
	}
	return i
}

func CheckForLeadingZeroes(hashString string, numZeroes int) bool {
	const zero = "0"
	for i := 0; i < numZeroes; i++ {
		hash := fmt.Sprintf("%x", hashString)
		if hash[i] != zero[0] {
			return false
		}

	}
	return true
}

func RunPuzzle2(secret string) int {
	i := 0
	hash := GetSumFromSecretAndAnswer(secret, strconv.Itoa(i))
	for failed := true; failed; failed = !CheckForLeadingZeroes(hash, 6) {
		i++
		hash = GetSumFromSecretAndAnswer(secret, strconv.Itoa(i))
	}
	return i
}
