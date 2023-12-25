package main

import (
	"fmt"
	. "github.com/iondodon/stream/stream"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}

	list, err := ToStream(list).
		Filter(even).
		Apply(multiplyTo10).
		Peek(printToConsole).
		ToSlice()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nResulting slice:", list)
}

func multiplyTo10(i int) (int, error) {
	return i * 10, nil
}

func printToConsole(e int) error {
	fmt.Print(e, " ")
	return nil
}

func even(e int) (bool, error) {
	if e%2 == 0 {
		return true, nil
	}
	return false, nil
}
