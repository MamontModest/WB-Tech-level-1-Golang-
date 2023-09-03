package main

import "fmt"

var arrStrings = []string{"cat", "cat", "dog", "cat", "tree", "Tree"}

func main() {
	set := ArrToSet(arrStrings)
	for i, _ := range set {
		fmt.Print(i, " ")
	}
}

// ArrToSet превращает массив строк в множество строк, регистрозависимая
func ArrToSet(arrS []string) map[string]bool {
	result := make(map[string]bool)
	for _, v := range arrS {
		result[v] = true
	}
	return result
}
