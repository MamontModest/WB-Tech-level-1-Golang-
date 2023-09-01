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
		go func(num int) {
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println(num * num)
			wg.Done()
		}(num)
	}
	wg.Wait()
}
