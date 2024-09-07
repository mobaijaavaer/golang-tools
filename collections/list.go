package collections

import "github.com/mobaijaavaer/golang-tools/functions"

// List 是一个泛型接口，定义了对列表数据结构的操作。
type List[T comparable] interface {
	Add(item T)
	AddAll(items []T)
	AddIf(item T, predicate functions.Predicate[T])
	Clear()
	Contains(item T) bool
	ContainsAll(items []T) bool
	Distinct()
	Get(index int) T
	GetData() []T
	IsEmpty() bool
	Remove(item T) bool
	Size() int
	ToString() string
	Sort(mapper functions.Function[T, int])
	NewEmpty() List[T]
}

// Filter 根据给定的谓词函数过滤 List 中的元素，并返回新的 List。
func Filter[I comparable](source List[I], filter functions.Predicate[I]) List[I] {
	res := source.NewEmpty()
	for _, item := range source.GetData() {
		if filter(item) {
			res.Add(item)
		}
	}
	return res
}

// Reduce 对 List 中的元素进行归约，根据给定的初值和归约函数返回最终结果。
func Reduce[I comparable, R comparable](source List[I], initialValue R, reducer functions.BiFunction[R, I, R]) R {
	res := initialValue
	for _, item := range source.GetData() {
		res = reducer(res, item)
	}
	return res
}

// ToMap 将 List 中的每个元素通过给定的映射函数转换为键值对，返回一个映射表。
func ToMap[I comparable, K comparable, V comparable](source List[I], mapper func(item I) (K, V)) map[K]V {
	res := make(map[K]V)
	for _, item := range source.GetData() {
		k, v := mapper(item)
		res[k] = v
	}
	return res
}

// GroupingBy 根据给定的映射函数对 List 中的元素进行分组，并返回分组结果。
func GroupingBy[T comparable, R comparable](source List[T], mapper functions.Function[T, R]) map[R]List[T] {
	res := make(map[R]List[T])
	for _, item := range source.GetData() {
		key := mapper(item)
		if _, ok := res[key]; !ok {
			res[key] = source.NewEmpty()
		}
		res[key].Add(item)
	}
	return res
}
