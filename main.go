package main

import "fmt"

type consumerFunc[T any] func(T)
type functionFunc[T any] func(T) T

type consumer interface {
	consume(elem interface{})
}

type function interface {
	apply(interface{}) interface{}
}

func (f consumerFunc[T]) consume(elem T) {
	f(elem)
}

func (f functionFunc[T]) apply(elem T) T {
	return f(elem)
}

type intStream struct {
	collection []int
	actions    []interface{}
}

func (s *intStream) do(action interface{}) *intStream {
	s.actions = append(s.actions, action)
	return s
}

func (s *intStream) toSlice() []int {
	for _, action := range s.actions {
		if action, ok := action.(consumer); ok {
			for _, elem := range s.collection {
				action.consume(elem)
			}
		}
		if a, ok := action.(function); ok {
			for index, elem := range s.collection {
				res := a.apply(elem)
				s.collection[index] = res.(int)
			}
		}
	}
	return s.collection
}

func toStream(collection []int) *intStream {
	return &intStream{
		collection: collection,
		actions:    nil,
	}
}

func main() {
	list := []int{1, 2, 3}

	s := toStream(list).
		do(func(i int) int { return i * 2 }).
		do(func(e int) int { fmt.Print(e, " "); return e }).
		toSlice()

	fmt.Println("\nResulting slice:", s)
}
