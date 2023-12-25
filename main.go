package main

import "fmt"

type action[T any] struct {
	actionType string
}

type peeker[T any] interface {
	peek(elem T)
}

type consumerFunc[T any] func(T) T

func (c consumerFunc[T]) peek(elem T) {
	c(elem)
}

type stream[T any] struct {
	collection []T

	actions []action[T]

	peekers   []peeker[T]
	peekIndex int
}

func (s *stream[T]) peek(actionFunc consumerFunc[T]) *stream[T] {
	s.actions = append(s.actions, action[T]{actionType: "peek"})
	s.peekers = append(s.peekers, actionFunc)
	return s
}

func (s *stream[T]) toSlice() []T {
	for _, a := range s.actions {
		if a.actionType == "peek" {
			s.peekIndex = s.peekIndex + 1
			if p, ok := s.peekers[s.peekIndex].(peeker[T]); ok {
				for _, elem := range s.collection {
					p.peek(elem)
				}
			}
		}
	}
	return s.collection
}

func toStream[T any](collection []T) *stream[T] {
	return &stream[T]{
		collection: collection,

		actions: make([]action[T], 0),

		peekers:   make([]peeker[T], 0),
		peekIndex: -1,
	}
}

func main() {
	list := []int{1, 2, 3}

	s := toStream(list).
		peek(func(e int) int { fmt.Print(e, " "); return e }).
		toSlice()

	fmt.Println("\nResulting slice:", s)
}
