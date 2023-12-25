package main

import (
	"fmt"
	. "github.com/iondodon/stream/stream"
)

func main() {
	list := []int{1, 2, 3}

	s := ToStream(list).
		Peek(printToConsole).
		Filter(even).
		ToSlice()

	fmt.Println("\nResulting slice:", s)
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
