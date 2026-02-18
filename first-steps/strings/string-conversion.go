package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int64 = 10
	ageStr := string(age)             // This will not convert the integer to a string representation of the number
	println("Age as string:", ageStr) // This will print a non-readable character
	ageStr = fmt.Sprintf("%d", age)   // Proper way to convert int to string
	println("Age as string:", ageStr) // This will print "Age as string: 10"
	ageStr = strconv.FormatInt(age, 10)
	println("Age as string:", ageStr) // This will print "Age as string: 10"
}
