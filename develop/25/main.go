package main

import (
	"fmt"
	"time"
)

func main() {
	startTime1 := time.Now()
	fmt.Println(startTime1)
	sleep1(5 * time.Second)

	endTime1 := time.Now()
	fmt.Println(endTime1)

	startTime2 := time.Now()
	fmt.Println(startTime2)
	sleep2(5 * time.Second)

	endTime2 := time.Now()
	fmt.Println(endTime2)
}

func sleep1(t time.Duration) {
	endTime := time.Now().Add(t)
	for {
		if endTime.Before(time.Now()) {
			return
		}
	}
}

func sleep2(t time.Duration) {
	<-time.After(t)
}
