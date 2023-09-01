package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var n int
	if err := input(&n); err != nil {
		fmt.Println("Вы ввели некорректное число")
		os.Exit(-1)
	}
	//канал при для успешного завершения
	osSigCh := make(chan os.Signal, 1)
	signal.Notify(osSigCh, os.Interrupt)
	c := make(chan string, 10)

	//Для синхронизации при завершении
	wg := sync.WaitGroup{}
	wg.Add(n)

	flag := false

	//Создаём горутины
	for i := 0; i < n; i++ {
		go func(i int) {
			for {
				select {
				//В Случае если можем вывести значение из канала
				case s := <-c:
					fmt.Printf("worker %d, %s \n", i, s)
					time.Sleep(3 * time.Second)
				default:
					//случай если ввели ctrl+c
					if flag {
						wg.Done()
						return
					}
					time.Sleep(1)
				}
			}
		}(i)
	}
out:
	for {
		select {
		case <-osSigCh:
			log.Println("CTRL+C interrupt by user, wait while goroutines end")
			flag = true
			break out
		default:
			c <- fmt.Sprintf("now is %s", time.Now().String())
			time.Sleep(2 * time.Second)
		}
	}
	wg.Wait()
}

// Input ввод числа с консоли
func input(n *int) error {
	fmt.Println("Write count of workers")
	_, err := fmt.Scanln(n)
	return err
}
