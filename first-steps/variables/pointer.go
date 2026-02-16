package main

import "fmt"

func DemoPointer() {
	fmt.Println("== Pointer demo ==")
	x := 42
	px := &x
	fmt.Println("value via pointer:", *px)
	*px = 100
	fmt.Println("x after change:", x)
}
