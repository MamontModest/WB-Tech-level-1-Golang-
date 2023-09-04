package main

import (
	"fmt"
)

func main() {
	var arr = []int{4, -5, 8, 3, 6, 11, -1, 13}
	fmt.Printf("не отсортированный массив %v \n", arr)
	quickSort(arr)
	fmt.Printf("отсортированный массив %v", arr)
}

// Функция изменяет массив и сортирует его
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	//опорный элемент
	supportElem := arr[len(arr)-1]

	i, j := 0, len(arr)-1
	//Ставим все элементы меньше либо равны опорному слева, остальные справа
	for i != j {
		if arr[i] <= supportElem && arr[j] > supportElem {
			j -= 1
		} else if arr[i] > supportElem && arr[j] > supportElem {
			j -= 1
		} else if arr[i] <= supportElem && arr[j] <= supportElem {
			i += 1
		} else if arr[i] > supportElem && arr[j] <= supportElem {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}
	//Если все элементы меньше опорного слева от i и наибольший элемент не в конце массива
	if arr[i] <= supportElem && i != len(arr)-1 {
		quickSort(arr[:i+1])
		quickSort(arr[i+1:])
	} else {
		quickSort(arr[:i])
		quickSort(arr[i:])
	}

}
