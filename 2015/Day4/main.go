package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	fmt.Println(GetSumFromSecretAndAnswer("abcdef", "609043"))
	fmt.Println(CheckFirst5Characters(GetSumFromSecretAndAnswer("abcdef", "609043")))
}

func GetSumFromSecretAndAnswer(secret, answer string) string {
	hash := md5.New()
	io.WriteString(hash, secret)
	io.WriteString(hash, answer)
	fmt.Printf("%x", hash.Sum(nil))
	return string(hash.Sum(nil))
}

func CheckFirst5Characters(hashString string) bool {
	const fiveZeroes = "00000"
	fmt.Println(hashString)
	for i := 0; i < 5; i++ {
		hash := fmt.Sprintf("%x", hashString)
		if hash[i] != fiveZeroes[i] {
			fmt.Println(hashString[i])
			return false
		}

	}
	return true
}
