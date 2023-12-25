package main

import (
	"fmt"
	. "github.com/iondodon/stream/stream"
)

func main() {
	list := []int{1, 2, 3}

	s := ToStream(list).
		Filter(even).
		Apply(multiplyTo10).
		Peek(printToConsole).
		ToSlice()

	fmt.Println("\nResulting slice:", s)
}

func multiplyTo10(i int) int {
	return i * 10
}

func printToConsole(e int) {
	fmt.Print(e, " ")
}

func even(e int) bool {
	if e%2 == 0 {
		return true
	}
	return false
}
