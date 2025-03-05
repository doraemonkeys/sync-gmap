package syncgmap

import "sync"

// SyncMap is a wrapper around sync.Map that is safe for concurrent use
// by multiple goroutines without additional locking or coordination.
// Loads, stores, and deletes run in amortized constant time.
//
// The SyncMap type is optimized for two common use cases: (1) when the entry for a given
// key is only ever written once but read many times, as in caches that only grow,
// or (2) when multiple goroutines read, write, and overwrite entries for disjoint
// sets of keys. In these two cases, use of a SyncMap may significantly reduce lock
// contention compared to a Go map paired with a separate [Mutex] or [RWMutex].
//
// The zero SyncMap is empty and ready for use. A SyncMap must not be copied after first use.
type SyncMap[K comparable, V any] struct {
	// sync.Map is exported for flexibility, so you can still
	// use it if required
	sync.Map
}

// func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
// 	return &SyncMap[K, V]{}
// }

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	result, ok := m.Map.Load(key)
	if ok {
		return result.(V), true
	}

	return *new(V), false
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	result, ok := m.Map.LoadOrStore(key, value)
	if ok {
		return result.(V), true
	}

	return value, false
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	item, ok := m.Map.LoadAndDelete(key)

	if ok {
		return item.(V), true
	}

	return *new(V), false
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func CompareAndDelete[K comparable, V comparable](m *SyncMap[K, V], key K, old V) (deleted bool) {
	return m.Map.CompareAndDelete(key, old)
}

func CompareAndSwap[K comparable, V comparable](m *SyncMap[K, V], key K, old, new V) (swapped bool) {
	return m.Map.CompareAndSwap(key, old, new)
}
