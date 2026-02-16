package main

import "fmt"

func DemoMap() {
	fmt.Println("== Map demo ==")
	m := map[string]int{"a": 1}
	m["b"] = 2
	v, ok := m["a"]
	fmt.Println("map:", m, "a value:", v, "found:", ok)
	delete(m, "a")
	fmt.Println("after delete:", m)
}
