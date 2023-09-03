package main

import "fmt"

func main() {
	a := 16
	b := 69

	swapIntBasic(&a, &b)
	fmt.Println(fmt.Sprintf("a : %d, b : %d", a, b))

	swapIntXOR(&a, &b)
	fmt.Println(fmt.Sprintf("a : %d, b : %d", a, b))
}

// Первый способ
func swapIntBasic(i, j *int) {
	*i, *j = *j, *i
}

// Второй способ
func swapIntXOR(i, j *int) {
	*i = *i ^ *j
	*j = *i ^ *j
	*i = *i ^ *j
}
