package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int64

	//Первый способ
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go worker(wg, &value)
	go worker(wg, &value)
	go worker(wg, &value)
	go worker(wg, &value)
	wg.Wait()
	fmt.Println(value)

	//Второй способ
	var valueS int
	wgS := &sync.WaitGroup{}
	mt := &sync.Mutex{}
	wgS.Add(4)
	go workerS(wgS, &valueS, mt)
	go workerS(wgS, &valueS, mt)
	go workerS(wgS, &valueS, mt)
	go workerS(wgS, &valueS, mt)
	wgS.Wait()
	fmt.Println(valueS)
}

// С помощью атомарных операций
func worker(wg *sync.WaitGroup, value *int64) {
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(value, 1)
	}
	wg.Done()

}

// С помощью mutex
func workerS(wg *sync.WaitGroup, value *int, mutex *sync.Mutex) {
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		*value++
		mutex.Unlock()
	}
	wg.Done()

}
