package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	endWG()
	time.Sleep(1 * time.Second)
	endChan()
	time.Sleep(1 * time.Second)
	endContext()
}

func endWG() {
	wg := &sync.WaitGroup{}
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("Потоки с завершением с помощью wg стартанули")
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(i int, rd *rand.Rand) {
			k := 0
			for k < 20 {
				n := rd.Uint64()
				if n > 10 && n%20 == 1 {
					k += 1
					fmt.Printf("goroutine %d find %d\n", i, n)
					time.Sleep(time.Millisecond * 100)
				}
			}
			wg.Done()
		}(i, rd)
	}
	wg.Wait()
	fmt.Println("Потоки с завершением с помощью wg закончились")
}

func endChan() {
	ch := make(chan struct{})
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("Потоки с завершением с помощью closeChan стартанули")
	for i := 0; i < 20; i++ {
		go func(i int, ch chan struct{}, rd *rand.Rand) {
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						fmt.Printf("goroutine %d dead \n", i)
						return
					}
				default:
					n := rd.Uint64()
					if n > 10 && n%20 == 1 {
						fmt.Printf("goroutine %d find %d\n", i, n)
						time.Sleep(time.Millisecond * 100)
					}
				}
			}
		}(i, ch, rd)
	}

	time.Sleep(5 * time.Second)
	close(ch)
	fmt.Println("Потоки с завершением с помощью closeChan закончились")
}

func endContext() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	fmt.Println("Поток с завершением с помощью context стартанули")
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	ch := make(chan bool)
	go func(i int, ctx context.Context, rd *rand.Rand, ch chan bool) {
		for {

			select {
			case <-ctx.Done():
				fmt.Printf("goroutine %d dead \n", i)
				ch <- true
				return
			default:
				n := rd.Uint64()
				if n > 10 && n%20 == 1 {
					fmt.Printf("goroutine %d find %d\n", i, n)
					time.Sleep(time.Millisecond * 100)
				}
			}
		}
	}(1, ctx, rd, ch)
	<-ch
	fmt.Println("Поток с завершением с помощью context закончились")

}
