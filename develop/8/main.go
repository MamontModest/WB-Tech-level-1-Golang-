package main

import (
	"errors"
	"fmt"
)

func main() {
	var n int64
	n = 15
	n = CircleReplace(n, 1, 1)  // 15
	n = CircleReplace(n, 1, 0)  // 14
	n = CircleReplace(n, 10, 1) //526
	n = CircleReplace(n, 2, 0)  // 524
	n = CircleReplace(n, 12, 1) // 2572
	//Чтение и с клавиатуры числа и изменение битов
	/*
		var n int64
		if err := input(&n); err != nil {
			fmt.Println("Вы ввели некорректное число")
			os.Exit(-1)
		}
		for {
			index, bit, err := inputBitsInfo()
			if err != nil{
				fmt.Println(err.Error())
				continue
			}
			n = CircleReplace(n, index, bit)
		}
	*/
}

func input(n *int64) error {
	fmt.Println("Введите число")
	_, err := fmt.Scan(n)
	return err
}

func CircleReplace(n int64, index, bit int) int64 {
	value, err := checkBit(n, index)
	if err != nil {
		fmt.Println(err.Error())
		return n
	}
	n = replaceBit(n, value, bit, index)
	fmt.Println(fmt.Sprintf("Новое значение n = %d", n))
	return n
}

func replaceBit(n int64, oldValue, newValue, index int) int64 {
	if oldValue == newValue {
		return n
	} else {
		if newValue == 1 {
			n += 1 << (index - 1)
		} else {
			n -= 1 << (index - 1)
		}
	}
	return n
}

func checkBit(n int64, i int) (int, error) {
	if i < 0 || i > 64 {
		return 0, errors.New("wrong position of bit")
	}
	return int((n >> (i - 1)) % 2), nil
}
func inputBitsInfo() (int, int, error) {
	var index, bit int
	fmt.Println("Введите какой бит нужно установить")
	_, err := fmt.Scan(&index)
	if err != nil {
		return 0, 0, errors.New("вы ввели некорректное число, попробуем снова")
	}
	fmt.Println("Введите значение для бита, 0 или 1")
	_, err = fmt.Scan(&bit)
	if err != nil {
		fmt.Println("Вы ввели некорректное число, попробуем снова")
		return 0, 0, errors.New("вы ввели некорректное число, попробуем снова")
	}
	return index, bit, nil
}
