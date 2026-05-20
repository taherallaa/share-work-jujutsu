package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	fmt.Println(add(1, 2))
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func divide(x, y int) int {
	return x / y
}

func multiple(x, y int) int {
	fmt.Println("I'm multiple function")
	return x * y
}

func foo() {
	fmt.Println("foo function")
}

func bar() {
	fmt.Println("bar function")
}
