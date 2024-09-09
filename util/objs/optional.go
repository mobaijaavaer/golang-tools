package objs

import "github.com/mobaijaavaer/golang-tools/functions"

var MEPTY = Optional[any]{value: nil}

type Optional[T any] struct {
	value *T
}

func Of[T any](value *T) *Optional[T] {
	if nil == value {
		panic("value is nil")
	}
	return &Optional[T]{value: value}
}

func OfNullable[T any](value *T) *Optional[T] {
	return &Optional[T]{value: value}
}

func (o *Optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o *Optional[T]) Get() T {
	if !o.IsPresent() {
		panic("value is nil")
	}
	return *o.value
}

func (o *Optional[T]) ifPresent(consumer functions.Consumer[T]) {
	if o.IsPresent() {
		consumer(o.Get())
	}
}

func (o *Optional[T]) filter(predicate functions.Predicate[T]) *Optional[T] {
	if !o.IsPresent() {
		return o
	}
	if predicate(o.Get()) {
		return o
	}
	return &Optional[T]{}
}

func (o *Optional[T]) Map(mapper functions.Function[T, *T]) *Optional[T] {
	if !o.IsPresent() {
		return &Optional[T]{}
	}
	return Of(mapper(o.Get()))
}

func (o *Optional[T]) OrElse(other T) T {
	if o.IsPresent() {
		return o.Get()
	}
	return other
}

func (o *Optional[T]) OrElseGet(supplier functions.Supplier[T]) T {
	return o.OrElse(supplier())
}
