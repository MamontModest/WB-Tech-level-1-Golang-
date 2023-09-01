package main

import (
	"fmt"
	"sync"
)

var array = []int{2, 4, 6, 8, 10}

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(len(array))
	for _, num := range array {
		// создаем анонимную функцию с аргументом num , которая будет выводить квадраты чисел
		go func(num int) {
			//Закрываем мьютекс
			mutex.Lock()
			//открываем мьютекс по выходу из функции
			defer mutex.Unlock()
			//выводим квадрат числа
			fmt.Println(num * num)
			//говорим что goroutine успешно завершилась
			wg.Done()
		}(num)
	}
	wg.Wait()
}
