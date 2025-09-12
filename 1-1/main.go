package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age  int
}

func (h *Human) Dance() string {
	return "Human with name " + h.name + " is dancing "
}

type Action struct {
	Human
	duration int
}

func (a *Action) DanceWithDuration() string {
	return a.Dance() + "in " + strconv.Itoa(a.duration) + "seconds"
}

func main() {
	a := &Action{
		Human: Human{
			name: "Artem",
			age:  22,
		},
		duration: 10,
	}
	fmt.Println(a.DanceWithDuration())
}
