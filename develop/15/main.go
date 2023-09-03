package main

import (
	"fmt"
	"strings"
)

/*
Здесь происходит утечка памяти, строка justString ссылается на тот-же объект памяти, что и hugeString,
но использует при этом всего 100 его символов => символы , которые не используются в justString не могут быть удаленны,
а могут лишь быть перезаписаны. По итогу мы тратим просто так память на хранение HugeString, которую уже не используем.
Так-же переменная justString глобальная, что является по себе плохой практикой и может привести к багу.
Например, если мы планируем вызывать функцию someFunc не один раз, мы можем не каждый раз генерировать переменную строку,
а создать ее единожды как константу и дальше использовать.

Var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/

// Создаём строку большой длины
func createHugeString(i int) string {
	builder := strings.Builder{}
	builder.Grow((i + 16) / 17)
	for j := 0; j < i; j = j + 17 {
		builder.WriteString("niceConcatenation")
	}
	return builder.String()
}

func someFunc(s *string) {
	v := createHugeString(1 << 10)
	//Не передаём тот-же участок памяти, а выделяем новый
	*s = strings.Clone(v)[:100]
}

func main() {
	var justString string
	someFunc(&justString)
	fmt.Println(justString)
}
