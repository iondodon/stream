package main

import "fmt"

type Action[T any] interface {
	Do(elem T) T
}

type consumer[T any] func(T) T
type producer[T any] func() T

func (c consumer[T]) Do(elem T) T {
	return c(elem)
}

func (p producer[T]) Do(elem T) T {
	return p()
}

type stream[T any] struct {
	collection []T
	actions    []Action[T]
}

func (s *stream[T]) peek(action consumer[T]) *stream[T] {
	s.actions = append(s.actions, action)
	return s
}

func (s *stream[T]) toSlice() []T {
	var slice []T
	for _, elem := range s.collection {
		for _, action := range s.actions {
			result := action.Do(elem)
			slice = append(slice, result)
		}
	}
	return slice
}

func toStream[T any](collection []T) *stream[T] {
	return &stream[T]{
		collection: collection,
		actions:    nil,
	}
}

func main() {
	list := []int{1, 2, 3}

	s := toStream(list).
		peek(func(e int) int {
			fmt.Print(e, " ")
			return e
		}).
		toSlice()

	fmt.Println("\nResulting slice:", s)
}
