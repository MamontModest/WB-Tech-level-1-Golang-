package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var arr = []int{2, 4, 6, 8, 10}

func main() {
	//sum := SumFirst(array)
	sum := SumSecond(arr)
	fmt.Println(sum)
}

// SumFirst способ подсчёта суммы номер один
func SumFirst(arr []int) int {
	sum := 0
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(len(arr))
	for _, num := range arr {
		//Создаем анонимную функцию с аргументом num , которая будет выводить квадраты чисел
		go func(num int) {
			//Закрываем мьютекс
			mutex.Lock()
			//Открываем мьютекс по выходу из функции
			defer mutex.Unlock()
			//Увеличиваем переменную
			sum += num * num
			//говорим что goroutine успешно завершилась
			wg.Done()
		}(num)
	}
	wg.Wait()
	return sum
}

// SumSecond способ подсчёта суммы номер два с атомарными операциями
func SumSecond(arr []int) int {
	var sum int64
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, num := range arr {
		//Создаем анонимную функцию с аргументом num которая будет выводить квадраты чисел
		go func(num int) {
			//Атомарно увеличиваем переменную sum
			atomic.AddInt64(&sum, int64(num*num))
			//говорим что goroutine успешно завершилась
			wg.Done()
		}(num)
	}
	wg.Wait()
	return int(sum)
}
