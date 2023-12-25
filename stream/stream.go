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

	appliers     []Applier[T]
	applierIndex int
}

func (s *Stream[T]) Peek(consumerFunc ConsumerFunc[T]) *Stream[T] {
	s.actions = append(s.actions, Action[T]{ActionType: PeekAction})
	s.peekers = append(s.peekers, consumerFunc)
	return s
}

func (s *Stream[T]) Filter(filterFunc PredicateFunc[T]) *Stream[T] {
	s.actions = append(s.actions, Action[T]{ActionType: FilterAction})
	s.filters = append(s.filters, filterFunc)
	return s
}

func (s *Stream[T]) Apply(functionFunc FunctionFunc[T]) *Stream[T] {
	s.actions = append(s.actions, Action[T]{ActionType: ApplyAction})
	s.appliers = append(s.appliers, functionFunc)
	return s
}

func (s *Stream[T]) ToSlice() []T {
	for _, a := range s.actions {
		switch a.ActionType {
		case PeekAction:
			s.doPeek()
		case FilterAction:
			s.doFilter()
		case ApplyAction:
			s.doApply()
		default:
			panic("unrecognized action")
		}
	}
	return s.collection
}

func (s *Stream[T]) doFilter() {
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

func (s *Stream[T]) doPeek() {
	s.peekIndex = s.peekIndex + 1
	if p, ok := s.peekers[s.peekIndex].(Peeker[T]); ok {
		for _, elem := range s.collection {
			p.Peek(elem)
		}
	}
}

func (s *Stream[T]) doApply() {
	s.applierIndex = s.applierIndex + 1
	if a, ok := s.appliers[s.applierIndex].(Applier[T]); ok {
		for index, _ := range s.collection {
			res := a.Apply(s.collection[index])
			s.collection[index] = res
		}
	}
}

func ToStream[T any](collection []T) *Stream[T] {
	return &Stream[T]{
		collection: collection,

		actions: make([]Action[T], 0),

		peekers:   make([]Peeker[T], 0),
		peekIndex: -1,

		filters:     make([]Filter[T], 0),
		filterIndex: -1,

		appliers:     make([]Applier[T], 0),
		applierIndex: -1,
	}
}
