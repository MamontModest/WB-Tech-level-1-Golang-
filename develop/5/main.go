package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	var n int
	if err := input(&n); err != nil {
		fmt.Println("Вы ввели некорректное число")
		os.Exit(-1)
	}
	//Создаем контекст
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	defer cancel()
	//Создаем генератор
	rd := rand.New(rand.NewSource(time.Now().Unix()))

	ch := make(chan int)
	endCh := make(chan interface{})

	wg := sync.WaitGroup{}
	wg.Add(1)

	//Запускаем goroutine
	go func() {
		for {
			select {
			//Смотрим завершиться ли контекст
			case <-endCh:
				wg.Done()
				return
			//Если есть значение в канале выводим его
			case s := <-ch:
				fmt.Println(fmt.Sprintf("worker print %d", s))
				time.Sleep(1 * time.Second)
			}
		}
	}()
out:
	for {
		select {
		case <-ctx.Done():
			endCh <- struct{}{}
			break out

		default:
			ch <- rd.Int()
			time.Sleep(1 * time.Second)
		}
	}
	//Ждём завершение goroutine
	wg.Wait()

}

// Input ввод числа с консоли
func input(n *int) error {
	fmt.Println("Write count of workers")
	_, err := fmt.Scanln(n)
	return err
}
