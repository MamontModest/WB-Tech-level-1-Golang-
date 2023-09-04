package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	fmt.Println(s1, " —  ", uniqLetters(s1))

	s2 := "abCdefAaf"
	fmt.Println(s2, " —  ", uniqLetters(s2))

	s3 := "aabcd"
	fmt.Println(s3, " — ", uniqLetters(s3))

	s4 := "aA"
	fmt.Println(s4, " — ", uniqLetters(s4))
}

func uniqLetters(s string) bool {
	hashMap := make(map[int32]bool)
	for _, v := range strings.ToLower(s) {
		if hashMap[v] {
			return false
		}
		hashMap[v] = true
	}
	return true
}
