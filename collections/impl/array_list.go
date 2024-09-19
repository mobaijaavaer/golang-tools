package impl

import (
	"fmt"
	"github.com/mobaijaavaer/golang-tools/collections"
	"github.com/mobaijaavaer/golang-tools/functions"
	"sort"
)

// ArrayList 是一个泛型列表结构，底层使用切片,用于存储任意类型的数据。
type ArrayList[T comparable] struct {
	data []T
}

// NewArrayList 创建并返回一个空的 ArrayList 实例。
func NewArrayList[T comparable]() *ArrayList[T] {
	return &ArrayList[T]{data: make([]T, 0)}
}

// NewArrayLists 根据给定的切片数据创建并返回一个 ArrayList 实例。
func NewArrayLists[T comparable](source []T) *ArrayList[T] {
	res := NewArrayList[T]()
	res.AddAll(source)
	return res
}

// Add 向 ArrayList 中添加一个元素。
func (a *ArrayList[T]) Add(item T) {
	a.data = append(a.data, item)
}

// Get 获取 ArrayList 中指定索引位置的元素。
func (a *ArrayList[T]) Get(index int) T {
	return a.data[index]
}

// AddIf 根据谓词条件向 ArrayList 中添加元素，只有当谓词返回 true 时才添加。
func (a *ArrayList[T]) AddIf(item T, predicate functions.Predicate[T]) {
	if !predicate(item) {
		return
	}
	a.data = append(a.data, item)
}

// AddAll 将给定切片中的所有元素添加到 ArrayList 中。
func (a *ArrayList[T]) AddAll(items []T) {
	if items == nil || len(items) == 0 {
		return
	}
	a.data = append(a.data, items...)
}

// Clear 清空 ArrayList 中的所有元素。
func (a *ArrayList[T]) Clear() {
	a.data = a.data[:0]
}

// Contains 检查 ArrayList 是否包含指定的元素。
func (a *ArrayList[T]) Contains(item T) bool {
	for _, v := range a.data {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsAll 检查 ArrayList 是否包含给定切片中的所有元素。
func (a *ArrayList[T]) ContainsAll(items []T) bool {
	for _, item := range items {
		if !a.Contains(item) {
			return false
		}
	}
	return true
}

// IsEmpty 检查 ArrayList 是否为空。
func (a *ArrayList[T]) IsEmpty() bool {
	return len(a.data) == 0
}

// Remove 从 ArrayList 中移除指定的元素，并返回移除是否成功。
func (a *ArrayList[T]) Remove(item T) bool {
	for i, v := range a.data {
		if v == item {
			a.data = append(a.data[:i], a.data[i+1:]...)
			return true
		}
	}
	return false
}

// Size 返回 ArrayList 中元素的数量。
func (a *ArrayList[T]) Size() int {
	return len(a.data)
}

// ToString 将 ArrayList 转换为字符串表示形式。
func (a *ArrayList[T]) ToString() string {
	return fmt.Sprintf("%v", a.data)
}

// Distinct 去除 ArrayList 中的重复元素。
func (a *ArrayList[T]) Distinct() {
	res := make([]T, 0)
	seen := make(map[T]bool)
	for _, value := range a.data {
		if !seen[value] {
			res = append(res, value)
			seen[value] = true
		}
	}
	a.data = res
}

func (a *ArrayList[T]) SortBy(comparator *collections.Comparator[T]) {
	sort.SliceStable(a.GetData(), func(i, j int) bool {
		return comparator.Compare(a.GetData()[i], a.GetData()[j]) < 0
	})
}

// Sort 对 ArrayList 中的元素进行排序，根据给定的映射函数返回排序后的 ArrayList。
func (a *ArrayList[T]) Sort(mapper functions.Function[T, int]) {
	sort.SliceStable(a.GetData(), func(i, j int) bool {
		return mapper(a.GetData()[i]) < mapper(a.GetData()[j])
	})
}

// NewEmpty 创建并返回一个新的空 ArrayList 实例。
func (a *ArrayList[T]) NewEmpty() collections.List[T] {
	return NewArrayList[T]()
}

// GetData 返回 ArrayList 中的数据切片。
func (this *ArrayList[T]) GetData() []T {
	return this.data
}
