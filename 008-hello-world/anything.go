package main

import "fmt"

func main() {
	fmt.Println("Hello everyone")
	foo()
	fmt.Println("something more here after foo call")
}

func foo() {
	fmt.Println("I'm in foo")
}

// control flow:
// 1. sequence
// 2. loop, iterative
// 3. conditional
