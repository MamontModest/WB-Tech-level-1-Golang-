package main

import (
	"fmt"
	"sort"
)

// Массив температур
var arr = []float32{1, -1, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

func main() {
	//создаем map, где группируем по температуре
	Temp := make(map[int][]float32)
	for _, v := range arr {
		//CloseT ближайшее значение для группы
		var CloseT int
		if v < 0 {
			CloseT = int(v) / 10 * 10
		} else {
			CloseT = int(v) / 10 * 10
		}
		Temp[CloseT] = append(Temp[CloseT], v)
	}
	//Создаем массив объектов, чтоб удобно было отсортировать по значениям температур групп
	windows := make([]Window, 0)
	for i, v := range Temp {
		windows = append(windows, Window{
			Point:        i,
			Temperatures: v,
		})
	}
	sort.Slice(windows, func(i, j int) bool {
		return windows[i].Point < windows[j].Point
	})
	for _, v := range windows {
		fmt.Println("Ближайшее значение для группы:", v.Point, ". Список температур :", v.Temperatures)
	}
}

type Window struct {
	Point        int
	Temperatures []float32
}
