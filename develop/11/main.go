package main

import (
	"fmt"
)

func main() {
	var arr1 = []int{1, 2, 3, 5, 9}
	var arr2 = []int{4, 5, 7, 10, -11, 9}
	fmt.Println(intersection(arr1, arr2))

	type point struct {
		x,
		y int
	}
	var arr3 = []point{{1, 2}, {2, 3}, {0, 0}}
	var arr4 = []point{{1, 2}, {2, 7}, {0, 0}, {5, 5}, {10, 1000}}
	fmt.Println(intersection(arr3, arr4))
}

// Intersection функция, которая реализует пересечение двух множеств
// В случае если все элементы в каждом из множеств уникальны
func intersection[T comparable](T1 []T, T2 []T) []T {
	//для красоты и понятности будем считать, что кол-во элементов списка T1 >= кол-во элементов списка T2
	if len(T1) < len(T2) {
		T1, T2 = T2, T1
	}
	//пересечение множеств в виде списка
	result := make([]T, 0)

	//дополнительная память для хранения T2 в виде множества
	hashMap := make(map[T]bool)

	//записываем все элементы T2 множество
	for _, v := range T2 {
		hashMap[v] = true
	}
	//проходимся по T1 проверяем есть ли этот элемент в T2
	for _, v := range T1 {
		if hashMap[v] {
			result = append(result, v)
		}
	}
	return result

}
