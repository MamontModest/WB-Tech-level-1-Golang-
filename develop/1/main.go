package main

import (
	"fmt"
	"time"
)

type Human struct {
	Male    bool
	Age     int16
	Name    string
	SurName string
}

func (h Human) hello() string {
	gender := "girl"
	if h.Male {
		gender = "man"
	}
	return fmt.Sprintf("Hello I am %s %d years old, you can calling me %s %s", gender, h.Age, h.Name, h.SurName)
}

func (h Human) NameSurName() string {
	return fmt.Sprintf("%s %s", h.Name, h.SurName)
}

type Action struct {
	ActionName       string
	Duration         time.Duration
	ActionComplexity string
	Human
}

func (a Action) Do() {
	s := a.hello()
	fmt.Printf("%s and for next %d hours,  i will %s. It going be %s.\n", s, int(a.Duration.Hours()), a.ActionName, a.ActionComplexity)
}

func (a Action) End() {
	fmt.Printf("%s %s end. Next time it will be easy for me.\n", a.NameSurName(), a.ActionName)
}

func main() {
	human := Human{
		Male:    true,
		Age:     24,
		Name:    "Ivan",
		SurName: "Kostirma",
	}
	action := Action{
		ActionName:       "work",
		Duration:         8 * time.Hour,
		ActionComplexity: "normal",
		Human:            human,
	}
	action.Do()

	action.End()
}
