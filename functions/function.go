package functions

// Function 是一个泛型函数类型，用于将一种类型的数据转换为另一种类型的数据。
type Function[E any, R any] func(E) R

// BiFunction 是一个泛型函数类型，用于接受两个参数并返回另一种类型的数据。
type BiFunction[E any, W any, R any] func(E, W) R

// Predicate 是一个泛型函数类型，用于对给定的数据进行判断，返回一个布尔值。
type Predicate[T any] func(T) bool

// BiConsumer 是一个泛型函数类型，用于接受两个参数并执行某些操作。
type BiConsumer[E any, W any] func(E, W)

// Consumer 是一个泛型函数类型，用于接受一个参数并执行某些操作。
type Consumer[T any] func(T)

type Supplier[T any] func() T
