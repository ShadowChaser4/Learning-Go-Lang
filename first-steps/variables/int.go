package main

import "fmt"

func DemoInt() {
	fmt.Println("== Int demo ==")
	var a int
	b := 5
	c := a + b
	fmt.Println("a:", a, "b:", b, "a+b:", c)
	fmt.Printf("Type of b: %T\n", b)
}
