package action

type ActionType string

const (
	ActionFilter = "filter"
	ActionPeek   = "peek"
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

type ConsumerFunc[T any] func(T)
type PredicateFunc[T any] func(T) bool

func (cf ConsumerFunc[T]) Peek(elem T) {
	cf(elem)
}

func (pf PredicateFunc[T]) Filter(elem T) bool {
	return pf(elem)
}
