package main

import "fmt"

func DemoBool() {
	fmt.Println("== Bool demo ==")
	t := true
	f := false
	fmt.Println("t && !f:", t && !f)
	fmt.Println("t || f:", t || f)
}
