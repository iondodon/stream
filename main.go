package main

import "fmt"

type action[T any] struct {
	actionType string
}

type peeker[T any] interface {
	peek(elem T)
}

type filter[T any] interface {
	filter(elem T) bool
}

type consumerFunc[T any] func(T)
type predicateFunc[T any] func(T) bool

func (cf consumerFunc[T]) peek(elem T) {
	cf(elem)
}

func (pf predicateFunc[T]) filter(elem T) bool {
	return pf(elem)
}

type stream[T any] struct {
	collection []T

	actions []action[T]

	peekers   []peeker[T]
	peekIndex int

	filters     []filter[T]
	filterIndex int
}

func (s *stream[T]) peek(consumerFunc consumerFunc[T]) *stream[T] {
	s.actions = append(s.actions, action[T]{actionType: "peek"})
	s.peekers = append(s.peekers, consumerFunc)
	return s
}

func (s *stream[T]) filter(filterFunc predicateFunc[T]) *stream[T] {
	s.actions = append(s.actions, action[T]{actionType: "filter"})
	s.filters = append(s.filters, filterFunc)
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
		} else if a.actionType == "filter" {
			s.filterIndex = s.filterIndex + 1
			if f, ok := s.filters[s.filterIndex].(filter[T]); ok {
				var res []T
				for _, elem := range s.collection {
					if f.filter(elem) {
						res = append(res, elem)
					}
				}
				s.collection = res
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

		filters:     make([]filter[T], 0),
		filterIndex: -1,
	}
}

func main() {
	list := []int{1, 2, 3}

	s := toStream(list).
		peek(toConsole).
		filter(even).
		toSlice()

	fmt.Println("\nResulting slice:", s)
}

func toConsole(e int) {
	fmt.Print(e, " ")
}

func even(e int) bool {
	if e%2 == 0 {
		return true
	}
	return false
}
