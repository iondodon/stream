package main

import (
	"errors"
	"fmt"
)

type StreamedSlice[T any] []T

type FunctionFunc[T any, R any] func(T) (R, error)

func trans[T any, R any](ss StreamedSlice[T], functionFunc FunctionFunc[T, R]) (StreamedSlice[R], error) {
	var newSlice = make(StreamedSlice[R], 0)

	for _, elem := range ss {
		res, err := functionFunc(elem)
		if err != nil {
			return nil, err
		}
		newSlice = append(newSlice, res)
	}

	return newSlice, nil
}

func (ss StreamedSlice[T]) collect() StreamedSlice[T] {
	return ss
}

func main() {
	s := StreamedSlice[int]([]int{1, 2, 3})

	res, err := trans(s, func(i int) (float32, error) {
		if i < 0 {
			return 0, errors.New("negative numbers not allowed")
		}
		return 2.21 + float32(i), nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(res.collect())
}
