package main

import (
	"fmt"
	"math/big"
)

func main() {
	num1 := big.NewInt(123)
	num2 := &big.Int{}

	num2.SetString("123123123123123123123123123123", 10) //"123"*10

	fmt.Printf("Наши числа: %s, %s \n", num1.String(), num2.String())

	fmt.Printf("Рузультат деления num2 на num1: %s \n", division(num2, num1).String())
	fmt.Printf("Рузультат умножения num2 на num1: %s \n", multiplication(num2, num1).String())
	fmt.Printf("Рузультат разницы num2 и num1: %s \n", subtraction(num2, num1).String())
	fmt.Printf("Рузультат суммы num2 и num1: %s \n", addition(num2, num1).String())
}

// деление
func division(num1, num2 *big.Int) *big.Int {
	result := &big.Int{}
	return result.Div(num1, num2)
}

// сложение
func addition(num1, num2 *big.Int) *big.Int {
	result := &big.Int{}
	return result.Add(num1, num2)
}

// вычитание
func subtraction(num1, num2 *big.Int) *big.Int {
	result := &big.Int{}
	return result.Sub(num1, num2)
}

// умножение
func multiplication(num1, num2 *big.Int) *big.Int {
	result := &big.Int{}
	return result.Mul(num1, num2)
}
