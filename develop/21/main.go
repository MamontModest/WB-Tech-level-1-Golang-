package main

import "fmt"

func main() {
	human := Human{weight: 65.6}
	fmt.Println("weight on earth: ", human.getBodyMassEarth())
	adapter := EarthToMoonAdapter{human}
	fmt.Println("weight on moon: ", adapter.getBodyMassMoon())
}

// Human Человек, который имеет переменную вес
type Human struct {
	weight float64
}

// BodyMassEarth Interface, который нужно адаптировать
type BodyMassEarth interface {
	getBodyMassEarth() float64
}

// BodyMassMoon интерфейс, к которому нужно адаптировать
type BodyMassMoon interface {
	getBodyMassMoon() float64
}

// EarthToMoonAdapter адаптер BodyMassEarth к BodyMassMoon
type EarthToMoonAdapter struct {
	BodyMassEarth
}

// Функция, которая возвращает массу тела на земле
func (h Human) getBodyMassEarth() float64 {
	return 9.81 * h.weight
}

// Функция реализующая адаптацию  BodyMassEarth к BodyMassMoon
func (b EarthToMoonAdapter) getBodyMassMoon() float64 {
	return b.getBodyMassEarth() / 9.8 * 1.62
}
