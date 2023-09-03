package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	indexValues := make(map[int][]string)
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	wg := &sync.WaitGroup{}

	//создаем мьютекс для конкурентной записи в map
	mutex := &sync.Mutex{}

	//wg для Завершения программы по выполнению всех goroutine
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go worker(i, rd, indexValues, wg, mutex)
	}
	wg.Wait()
	for _, v := range indexValues {
		//выводим , что получилось
		fmt.Println(strings.Join(v, ", "))
	}
}

func worker(w int, rd *rand.Rand, hashMap map[int][]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 0; i < 5; i++ {
		index := rd.Int() % 10
		//перед записью закрываем мьютекс
		mutex.Lock()
		hashMap[index] = append(hashMap[index], fmt.Sprintf("worker %d was here", w))
		//после записи открываем мьютекс
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
	wg.Done()
}
