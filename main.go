package main

import "fmt"

type Action[T any] interface {
	Do(elem T) T
}

type consumer[T any] func(T) T
type function[T any, R any] func(T) R

func (c consumer[T]) Do(elem T) T {
	return c(elem)
}

func (f function[T, R]) Do(elem T) R {
	return f(elem)
}

type stream[T any] struct {
	collection []T
	actions    []Action[T]
}

func (s *stream[T]) peek(action consumer[T]) *stream[T] {
	s.actions = append(s.actions, action)
	return s
}

func (s *stream[T]) transform(action function[T, T]) *stream[T] {
	s.actions = append(s.actions, action)
	return &stream[T]{
		collection: s.collection,
		actions:    s.actions,
	}
}

func (s *stream[T]) toSlice() []T {
	for _, action := range s.actions {
		var slice []T
		for _, elem := range s.collection {
			result := action.Do(elem)
			slice = append(slice, result)
		}
		s.collection = slice
	}
	return s.collection
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
		transform(func(i int) int { return i * 2 }).
		peek(func(e int) int { fmt.Print(e, " "); return e }).
		toSlice()

	fmt.Println("\nResulting slice:", s)
}
