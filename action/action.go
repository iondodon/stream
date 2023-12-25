package action

type ActionType string

const (
	FilterAction = "filter"
	PeekAction   = "peek"
	ApplyAction  = "apply"
)

type Action[T any] struct {
	ActionType string
}

type Peeker[T any] interface {
	Peek(elem T)
}

type Filter[T any] interface {
	Filter(elem T) bool
}

type Applier[T any] interface {
	Apply(elem T) T
}

type ConsumerFunc[T any] func(T)
type PredicateFunc[T any] func(T) bool
type FunctionFunc[T any] func(T) T

func (cf ConsumerFunc[T]) Peek(elem T) {
	cf(elem)
}

func (pf PredicateFunc[T]) Filter(elem T) bool {
	return pf(elem)
}

func (ff FunctionFunc[T]) Apply(elem T) T {
	return ff(elem)
}
