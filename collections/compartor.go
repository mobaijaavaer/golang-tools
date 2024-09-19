package collections

type Comparator[T any] struct {
	Comparing func(a T, b T) int
}

func (c *Comparator[T]) ThenComparing(compartor func(a T, b T) int) *Comparator[T] {
	return &Comparator[T]{Comparing: func(a T, b T) int {
		re := c.Comparing(a, b)
		if re == 0 {
			return compartor(a, b)
		}
		return re
	}}
}

func (c *Comparator[T]) Reverse() *Comparator[T] {
	return &Comparator[T]{Comparing: func(a T, b T) int {
		return -c.Comparing(a, b)
	}}
}

func (c *Comparator[T]) Compare(a T, b T) int {
	return c.Comparing(a, b)
}
