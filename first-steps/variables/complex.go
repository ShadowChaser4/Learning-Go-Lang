package main

import "fmt"

func DemoComplex() {
	complexNum := 3 + 4i
	fmt.Println("Complex number:", complexNum)
	fmt.Println("Real part:", real(complexNum))
	fmt.Println("Imaginary part:", imag(complexNum))
}
