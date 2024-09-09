package collections

type Queue[T comparable] interface {
	Offer(item *T) bool

	Poll() *T

	Peek() *T
}
