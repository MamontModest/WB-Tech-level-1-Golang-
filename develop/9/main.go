package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 5; i++ {
		//добавляют x в канал ch1
		go workerOne(ch1, rd)
		//добавляют x*2 в канал ch2
		go workerSecond(ch1, ch2)
	}
	//выводит x*2
	go workerThird(ch2)
	time.Sleep(time.Second * 10)
}

func workerOne(ch1 chan<- int, rd *rand.Rand) {
	for {
		ch1 <- rd.Int()
		time.Sleep(time.Second)
	}

}

func workerSecond(ch1 <-chan int, ch2 chan<- int) {
	for {
		ch2 <- 2 * <-ch1
		time.Sleep(time.Second)
	}
}

func workerThird(ch2 <-chan int) {
	for {
		fmt.Println(<-ch2)
	}
}
