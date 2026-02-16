## Variables and Constants in Go

In Go, variables and constants are fundamental building blocks for storing and manipulating data.

### Variables

A variable is a named storage location that can hold a value of a specific type. You can declare a variable using the `var` keyword, followed by the variable name and its type. For example:

```go
var age int = 30
```

You can also declare a variable without specifying its type, and Go will infer the type based on the assigned value:

```go
var name = "Alice"
```

### Constants

A constant is a named value that cannot be changed after it is assigned. Constants are declared using
the `const` keyword. For example:

```go
const Pi = 3.14
```

Constants can also be of any basic data type, such as strings, integers, and booleans:

```go
const Greeting = "Hello, World!"
const IsGoFun = true
```

### Short Variable Declaration

Go also provides a shorthand syntax for declaring and initializing variables using the `:=` operator. This is only allowed within functions. For example:

```go
func main() {
    message := "Welcome to Go!"
    fmt.Println(message)
}
```

In this example, the variable `message` is declared and initialized with the value "Welcome to Go!" using the short variable declaration syntax.

### Zero Values

When a variable is declared without an explicit initial value, it is assigned a zero value based on its type. For example:

```go
var count int    // count is initialized to 0
var isActive bool // isActive is initialized to false
var name string   // name is initialized to an empty string ""
```

### Conclusion

Understanding how to use variables and constants is essential for writing effective Go programs. Variables allow you to store and manipulate data, while constants provide a way to define values that should not change throughout the program.
By using these constructs appropriately, you can create more readable and maintainable code.
