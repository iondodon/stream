package action

type Type uint

const (
	FilterAction Type = iota
	PeekAction
	ApplyAction
)

type Action[T any] struct {
	ActionType Type
}

type Peeker[T any] interface {
	Peek(elem T) error
}

type Filter[T any] interface {
	Filter(elem T) (bool, error)
}

type Applier[T any] interface {
	Apply(elem T) (T, error)
}

type ConsumerFunc[T any] func(T) error
type PredicateFunc[T any] func(T) (bool, error)
type FunctionFunc[T any] func(T) (T, error)

func (cf ConsumerFunc[T]) Peek(elem T) error {
	return cf(elem)
}

func (pf PredicateFunc[T]) Filter(elem T) (bool, error) {
	return pf(elem)
}

func (ff FunctionFunc[T]) Apply(elem T) (T, error) {
	return ff(elem)
}
