package concurrent

import (
	"sync"
	"sync/atomic"
)

const MAX_CAPACITY = (1 << 31) - 1

type LinkedBlockingQueue[T any] struct {
	head *Node[T]

	tail *Node[T]

	size int

	count atomic.Int32

	putLock *sync.Mutex

	takeLock *sync.Mutex

	notEmpty *sync.Cond

	notFull *sync.Cond
}

type Node[T any] struct {
	item *T

	next *Node[T]
}

func NewLinkedBlockingQueue[T any](capacity int) *LinkedBlockingQueue[T] {
	if capacity < 0 || capacity > MAX_CAPACITY {
		panic("capacity must be greater than 0 and less than MAX_CAPACITY")
	}
	emptyNode := new(Node[T])
	res := &LinkedBlockingQueue[T]{
		head:     emptyNode,
		tail:     emptyNode,
		size:     capacity,
		count:    atomic.Int32{},
		putLock:  &sync.Mutex{},
		takeLock: &sync.Mutex{},
	}

	res.notFull = sync.NewCond(res.putLock)
	res.notEmpty = sync.NewCond(res.takeLock)
	return res
}

func NewDefaultLinkedBlockingQueue[T any]() *LinkedBlockingQueue[T] {
	return NewLinkedBlockingQueue[T](MAX_CAPACITY)
}

func (q *LinkedBlockingQueue[T]) Offer(item *T) bool {
	if item == nil {
		panic("item must not be nil")
	}
	n := &Node[T]{item: item}
	c := -1
	q.putLock.Lock()
	defer func() {
		q.putLock.Unlock()
		if c == 0 {
			q.takeLock.Lock()
			defer q.takeLock.Unlock()
			q.notEmpty.Signal()
		}
	}()
	for q.count.Load() == int32(q.size) {
		q.notFull.Wait()
	}
	q.tail.next = n
	q.tail = n
	c = int(q.count.Swap(q.count.Load() + 1))
	if c+1 < q.size {
		q.notFull.Signal()
	}
	return true
}

func (q *LinkedBlockingQueue[T]) Poll() *T {
	c := -1
	q.takeLock.Lock()
	defer func() {
		q.takeLock.Unlock()
		if c == q.size {
			q.putLock.Lock()
			defer q.putLock.Unlock()
			q.notFull.Signal()
		}
	}()
	for q.count.Load() == 0 {
		q.notEmpty.Wait()
	}
	x := q.dequeue()
	c = int(q.count.Swap(q.count.Load() - 1))
	if c > 1 {
		q.notEmpty.Signal()
	}
	return x
}

func (q *LinkedBlockingQueue[T]) Peek() *T {
	q.takeLock.Lock()
	defer q.takeLock.Unlock()
	if q.count.Load() == 0 {
		return nil
	}
	first := q.head.next
	if first == nil {
		return nil
	}
	return q.head.next.item
}

func (q *LinkedBlockingQueue[T]) dequeue() *T {
	h := q.head
	first := h.next
	h.next = h
	q.head = first
	x := first.item
	first.item = nil
	return x
}
