package main

import "fmt"

func DemoFloat() {
	fmt.Println("== Float demo ==")
	f := 3.14
	var f32 float32 = 1.5
	fmt.Println("f:", f, "f32:", f32)
	fmt.Printf("Sum (converted): %f\n", f+float64(f32))
}
