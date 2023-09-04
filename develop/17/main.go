package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 5, 5, 7, 10, 180}
	fmt.Println(binarySearch(arr, 5)) //3

	arr = []int{1, 2, 3, 4, 5, 9}
	fmt.Println(binarySearch(arr, 9)) //5

	arr = []int{1, 2, 3, 4, 5, 9, 9, 10}
	fmt.Println(binarySearch(arr, 9)) //6

	arr = []int{1, 2}
	fmt.Println(binarySearch(arr, 1)) //0

	arr = []int{1}
	fmt.Println(binarySearch(arr, 0)) //-1
}

// binarySearch ищет индекс самого левого элемента такого, что arr[index]<=elem и arr[index+1]>elem, если в массиве
// присутствует несколько index  таких , что arr[index] = elem , то он выведет index самого правого из них, (max(index))
// если elem < любого значения в массиве , то функция вернёт -1
func binarySearch(arr []int, elem int) int {
	if arr[0] > elem {
		return -1
	}
	l, r := 0, len(arr)-1
	for l != r {
		m := (l + r + 1) / 2
		if arr[m] <= elem {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
