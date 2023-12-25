package main

import (
	"fmt"
	. "github.com/iondodon/stream/stream"
)

func main() {
	list := []int{1, 2, 3}

	list, err := ToStream(list).
		Filter(func(e int) (bool, error) {
			if e%2 == 0 {
				return true, nil
			}
			return false, nil
		}).
		Apply(func(i int) (int, error) {
			return i * 10, nil
		}).
		Apply(func(i int) (int, error) {
			return i + 1, nil
		}).
		Peek(func(e int) error {
			fmt.Print(e, " ")
			return nil
		}).
		ToSlice()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nResulting slice:", list)
}
