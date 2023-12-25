package stream

import (
	. "github.com/iondodon/stream/action"
)

type Stream[T any] struct {
	collection []T

	actions []Action[T]

	peekers   []Peeker[T]
	peekIndex int

	filters     []Filter[T]
	filterIndex int
}

func (s *Stream[T]) Peek(consumerFunc ConsumerFunc[T]) *Stream[T] {
	s.actions = append(s.actions, Action[T]{ActionType: ActionPeek})
	s.peekers = append(s.peekers, consumerFunc)
	return s
}

func (s *Stream[T]) Filter(filterFunc PredicateFunc[T]) *Stream[T] {
	s.actions = append(s.actions, Action[T]{ActionType: ActionFilter})
	s.filters = append(s.filters, filterFunc)
	return s
}

func (s *Stream[T]) ToSlice() []T {
	for _, a := range s.actions {
		if a.ActionType == ActionPeek {
			s.peekIndex = s.peekIndex + 1
			if p, ok := s.peekers[s.peekIndex].(Peeker[T]); ok {
				for _, elem := range s.collection {
					p.Peek(elem)
				}
			}
		} else if a.ActionType == ActionFilter {
			s.filterIndex = s.filterIndex + 1
			if f, ok := s.filters[s.filterIndex].(Filter[T]); ok {
				var res []T
				for _, elem := range s.collection {
					if f.Filter(elem) {
						res = append(res, elem)
					}
				}
				s.collection = res
			}
		}
	}
	return s.collection
}

func ToStream[T any](collection []T) *Stream[T] {
	return &Stream[T]{
		collection: collection,

		actions: make([]Action[T], 0),

		peekers:   make([]Peeker[T], 0),
		peekIndex: -1,

		filters:     make([]Filter[T], 0),
		filterIndex: -1,
	}
}
