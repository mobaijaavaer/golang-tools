package collections

type Map[K comparable, V comparable] interface {
	Clear()

	ContainsKey(key K) bool

	Get(key K) V

	GetKeys() List[K]

	GetValues() List[V]

	IsEmpty() bool

	Put(key K, value V)

	PutAll(mapper Map[K, V])

	Remove(key K)

	Size() int

	GetOrDefault(key K, defaultValue V) V

	PutIfAbsent(key K, value V)

	Replace(key K, value V)
}
