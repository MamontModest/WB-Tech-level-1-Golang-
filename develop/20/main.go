package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "snow dog sun"
	fmt.Println(s1, " : ", reverseWords(s1))

	s2 := "снег собака солнце"
	fmt.Println(s2, " : ", reverseWords(s2))
}

//Разделяем строку по пробелам и меняем местами слова i с начала, с i с конца
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	n := len(words) - 1
	for i := 0; i < len(words)/2; i++ {
		words[i], words[n-i] = words[n-i], words[i]
	}
	return strings.Join(words, " ")
}
