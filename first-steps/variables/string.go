package main

import "fmt"

func DemoString() {
	fmt.Println("== String demo ==")
	s := "Hello"
	s2 := "World"
	fmt.Println(s + " " + s2)
	fmt.Printf("Length of %s: %d\n", s, len(s))
}
