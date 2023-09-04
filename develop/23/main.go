package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("before: %v\n", arr)

	newArr, err := removeIndex(arr, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("after: %v\n", newArr)

	//вернет ошибку
	newArr, err = removeIndex(arr, 8)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("after: %v\n", newArr)
}

//Возвращаем ошибку, если индекс за пределами массива
func removeIndex(arr []int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) {
		return nil, errors.New(fmt.Sprintf("index error. Your index %d, but index must be 0<=index<=%d", index, len(arr)-1))
	}
	return append(arr[:index], arr[index+1:]...), nil
}
