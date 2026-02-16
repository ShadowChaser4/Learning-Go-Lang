package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }

func (d Dog) Speak() string { return "Woof, I'm " + d.Name }

func DemoInterface() {
	fmt.Println("== Interface demo ==")
	var s Speaker = Dog{"Rex"}
	fmt.Println(s.Speak())
	switch v := s.(type) {
	case Dog:
		fmt.Println("It's a Dog:", v.Name)
	default:
		fmt.Println("Unknown type")
	}
}
