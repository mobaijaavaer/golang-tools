package impl

import (
	"github.com/mobaijaavaer/golang-tools/collections"
	"github.com/mobaijaavaer/golang-tools/functions"
)

type LinkList[T comparable] struct {
	head *Node[T]

	tail *Node[T]

	size int
}

type Node[T comparable] struct {
	data T

	next *Node[T]

	prev *Node[T]
}

func newNode[T comparable](prev *Node[T], data T, next *Node[T]) *Node[T] {
	return &Node[T]{data: data, next: next, prev: prev}
}

func NewLinkList[T comparable]() *LinkList[T] {
	return &LinkList[T]{}
}

func (l *LinkList[T]) Add(item T) {
	temp := l.tail
	node := newNode[T](temp, item, nil)
	if temp == nil {
		l.head = node
	} else {
		temp.next = node
	}
	l.tail = node
	l.size++
}

func (l *LinkList[T]) AddAll(items []T) {
	for _, item := range items {
		l.Add(item)
	}
}

func (l *LinkList[T]) AddIf(item T, predicate functions.Predicate[T]) {
	if predicate(item) {
		l.Add(item)
	}
}

func (l *LinkList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkList[T]) Contains(item T) bool {
	if l.head == nil {
		return false
	}
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.data == item {
			return true
		}
	}
	return false
}

func (l *LinkList[T]) ContainsAll(items []T) bool {
	for _, item := range items {
		if !l.Contains(item) {
			return false
		}
	}
	return true
}

func (l *LinkList[T]) Distinct() {
	for temp := l.head; temp != nil; temp = temp.next {
		for temp2 := temp.next; temp2 != nil; temp2 = temp2.next {
			if temp.data == temp2.data {
				temp2.prev.next = temp2.next
				if temp2.next != nil {
					temp2.next.prev = temp2.prev
				} else {
					l.tail = temp2.prev
				}
				l.size--
			}
		}
	}
}

func (l *LinkList[T]) Get(index int) T {
	if index < 0 || index > l.size {
		panic("index out of bounds")
	}
	//前序遍历 || 后序遍历
	if index < l.size>>1 {
		x := l.head
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x.data
	} else {
		x := l.tail
		for i := l.size - 1; i > index; i-- {
			x = x.prev
		}
		return x.data
	}
}

func (l *LinkList[T]) GetData() []T {
	data := make([]T, l.size)
	for i, temp := 0, l.head; temp != nil; temp = temp.next {
		data[i] = temp.data
		i++
	}
	return data
}

func (l *LinkList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkList[T]) Remove(item T) bool {
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.data == item {
			if temp.prev != nil {
				temp.prev.next = temp.next
			} else {
				l.head = temp.next
			}
			if temp.next != nil {
				temp.next.prev = temp.prev
			} else {
				l.tail = temp.prev
			}
			l.size--
			return true
		}
	}
	return false
}

func (l *LinkList[T]) Size() int {
	return l.size
}

func (l *LinkList[T]) ToString() string {
	return "nodes"
}

func (a *LinkList[T]) SortBy(comparator *collections.Comparator[T]) {
	//
}

// 链表非有序集合
func (l *LinkList[T]) Sort(mapper functions.Function[T, int]) {
	//TODO implement me
}

func (l *LinkList[T]) NewEmpty() collections.List[T] {
	return NewLinkList[T]()
}
