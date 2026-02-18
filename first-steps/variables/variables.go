package main

import "fmt"

func main() {
	DemoInt()
	DemoFloat()
	DemoBool()
	DemoString()
	DemoArraySlice()
	DemoMap()
	DemoStruct()
	DemoPointer()
	DemoInterface()
	DemoChannel()
	DemoComplex()

	name := "John"
	// taking input from user
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
