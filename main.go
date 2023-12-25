package main

import "fmt"

type StreamedSlice[T any] []T

type FunctionFunc[T any, R any] func(T) R

func trans[T any, R any](ss StreamedSlice[T], functionFunc FunctionFunc[T, R]) StreamedSlice[R] {
	var newSlice = make(StreamedSlice[R], 0)

	for _, elem := range ss {
		res := functionFunc(elem)
		newSlice = append(newSlice, res)
	}

	return newSlice
}

func (ss StreamedSlice[T]) collect() StreamedSlice[T] {
	return ss
}

func main() {
	s := StreamedSlice[int]([]int{1, 2, 3})

	res := trans(s, func(i int) float32 {
		return 2.21 + float32(i)
	}).collect()

	fmt.Println(res)
}
