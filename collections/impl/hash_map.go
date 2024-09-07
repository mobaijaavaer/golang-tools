package impl

import "github.com/mobaijaavaer/golang-tools/collections"

// HashMap 是一个键值对集合，它允许快速的查找和更新元素。
// 它使用了内建的map作为底层数据结构，提供了灵活的泛型定义。
// 它不保证元素的顺序，并且不允许重复的键。
type HashMap[K comparable, V comparable] struct {
	data map[K]V // 底层的map数据结构，存储键值对
}

// NewHashMap 创建并返回一个新的HashMap实例。
// 它初始化了底层的map，并返回一个指向该HashMap的指针。
func NewHashMap[K comparable, V comparable]() *HashMap[K, V] {
	return &HashMap[K, V]{data: make(map[K]V)}
}

// Clear 清空HashMap中的所有元素。
// 它通过重新创建底层map来达到清空的目的。
func (h *HashMap[K, V]) Clear() {
	h.data = make(map[K]V)
}

// ContainsKey 检查给定的键是否存在于HashMap中。
// 参数 key 是要检查的键。
// 返回值表示键是否存在于HashMap中。
func (h *HashMap[K, V]) ContainsKey(key K) bool {
	_, ok := h.data[key]
	return ok
}

// Get 获取与给定键相关联的值。
// 参数 key 是要查找的键。
// 返回值是与键相关联的值，如果键不存在则可能返回零值。
func (h *HashMap[K, V]) Get(key K) V {
	return h.data[key]
}

// GetKeys 获取HashMap中所有键的列表。
// 返回值是一个包含所有键的List。
func (h *HashMap[K, V]) GetKeys() collections.List[K] {
	keys := make([]K, len(h.data))
	i := 0
	for k := range h.data {
		keys[i] = k
		i++
	}
	return NewArrayLists(keys)
}

// GetOrDefault 获取与给定键相关联的值，如果键不存在则返回默认值。
// 参数 key 是要查找的键。
// 参数 defaultValue 是当键不存在时返回的默认值。
// 返回值是与键相关联的值或默认值。
func (h *HashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if h.ContainsKey(key) {
		return h.Get(key)
	}
	return defaultValue
}

// GetValues 获取HashMap中所有值的列表。
// 返回值是一个包含所有值的List。
func (h *HashMap[K, V]) GetValues() collections.List[V] {
	values := make([]V, len(h.data))
	i := 0
	for _, v := range h.data {
		values[i] = v
		i++
	}
	return NewArrayLists(values)
}

// IsEmpty 检查HashMap是否为空。
// 返回值表示HashMap是否为空。
func (h *HashMap[K, V]) IsEmpty() bool {
	return len(h.data) == 0
}

// Put 在HashMap中插入一个键值对。
// 参数 key 是要插入的键。
// 参数 value 是要插入的值。
// 该方法通过将键值对添加到底层map来实现插入。
func (h *HashMap[K, V]) Put(key K, value V) {
	h.data[key] = value
}

// PutAll 将另一个Map中的所有键值对放入HashMap中。
// 参数 mapper 是包含要放入键值对的Map。
// 该方法通过遍历另一个Map的键并将其值放入HashMap来实现。
func (h *HashMap[K, V]) PutAll(mapper collections.Map[K, V]) {
	for _, k := range mapper.GetKeys().GetData() {
		h.Put(k, mapper.Get(k))
	}
}

// PutIfAbsent 如果键不存在，则在HashMap中插入一个键值对。
// 参数 key 是要插入的键。
// 参数 value 是要插入的值。
// 该方法首先检查键是否已存在，如果不存在则插入。
func (h *HashMap[K, V]) PutIfAbsent(key K, value V) {
	if !h.ContainsKey(key) {
		h.Put(key, value)
	}
}

// Remove 从HashMap中移除与给定键相关联的元素。
// 参数 key 是要移除的键。
// 该方法通过从底层map中删除键来实现移除。
func (h *HashMap[K, V]) Remove(key K) {
	delete(h.data, key)
}

// Replace 替换HashMap中与给定键相关联的值。
// 参数 key 是要替换的键。
// 参数 value 是新的值。
// 该方法通过更新底层map中键对应的值来实现替换。
func (h *HashMap[K, V]) Replace(key K, value V) {
	h.data[key] = value
}

// Size 返回HashMap中键值对的数量。
// 返回值是HashMap中键值对的数量。
func (h *HashMap[K, V]) Size() int {
	return len(h.data)
}
