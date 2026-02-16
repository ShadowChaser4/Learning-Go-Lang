package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func DemoStruct() {
	fmt.Println("== Struct demo ==")
	p := Person{"Alice", 30}
	fmt.Println("Name:", p.Name, "Age:", p.Age)
}
