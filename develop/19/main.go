package main

import (
	"fmt"
)

func main() {
	fmt.Println("главрыба — ", reverseString("главрыба"))

	var s string
	for {
		input(&s)
		newString := reverseString(s)
		fmt.Println(newString)
	}
}

func input(s *string) {
	fmt.Println("Введите строку")
	fmt.Scan(s)
}

// Создаем слайс рун и меняем местами i и n-i элементы
func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes) - 1
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[n-i] = runes[n-i], runes[i]
	}
	return string(runes)
}
