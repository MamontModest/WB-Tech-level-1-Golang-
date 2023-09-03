package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	i := 6
	s := "String"
	f := true
	chn := make(chan int)
	hashMap := make(map[int]int)

	fmt.Println("First way")
	PrintTypeWithError(TypeOfCase(i))
	PrintTypeWithError(TypeOfCase(s))
	PrintTypeWithError(TypeOfCase(f))
	PrintTypeWithError(TypeOfCase(chn))
	PrintTypeWithError(TypeOfCase(hashMap))

	fmt.Println("Second  way")
	fmt.Println(TypeOfReflect(i))
	fmt.Println(TypeOfReflect(s))
	fmt.Println(TypeOfReflect(f))
	fmt.Println(TypeOfReflect(chn))
	fmt.Println(TypeOfReflect(hashMap))
}

func TypeOfReflect(value interface{}) reflect.Type {
	//Возвращаем тип
	return reflect.TypeOf(value)
}

func TypeOfCase(value interface{}) (reflect.Kind, error) {
	//Определяем тип с помощью кейса
	switch value.(type) {
	case int:
		return reflect.Int, nil
	case string:
		return reflect.String, nil
	case bool:
		return reflect.Bool, nil
	case chan interface{}:
		return reflect.Chan, nil
	case chan int:
		return reflect.Chan, nil
	case chan string:
		return reflect.Chan, nil
	case chan bool:
		return reflect.Chan, nil
	}
	return reflect.UnsafePointer, errors.New("undefined type")
}

func PrintTypeWithError(kind reflect.Kind, err error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(kind)
	}
}
