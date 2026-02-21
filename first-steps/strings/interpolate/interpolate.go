package main

import "fmt"

func main() {
	name := "Alice"
	greeting := fmt.Sprintf("Hello, %s!", name)

	fmt.Println(greeting, '\n')

	discount := 20.233
	message := "You are entitled to discount of"
	messageWithTwoDecimals := fmt.Sprintf("%s %.2f%%", message, discount)
	fmt.Printf("%s\n", messageWithTwoDecimals)
	messageWithOneDecimal := fmt.Sprintf("%s %.1f%%", message, discount)
	fmt.Printf("%s\n", messageWithOneDecimal)
}
