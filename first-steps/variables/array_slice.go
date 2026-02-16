package main

import "fmt"

func DemoArraySlice() {
	fmt.Println("== Array & Slice demo ==")
	arr := [3]int{1, 2, 3}
	sl := []int{4, 5}
	sl = append(sl, 6)
	fmt.Println("array:", arr)
	fmt.Println("slice:", sl)
	fmt.Println("slice[1:]:", sl[1:])
}
