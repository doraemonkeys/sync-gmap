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

// Load returns the value stored in the map for a key.
// The ok result indicates whether value was found in the map.
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	result, ok := m.Map.Load(key)
	if ok {
		return result.(V), true
	}

	return *new(V), false
}

// Store sets the value for a key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	result, ok := m.Map.LoadOrStore(key, value)
	if ok {
		return result.(V), true
	}

	return value, false
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	item, ok := m.Map.LoadAndDelete(key)

	if ok {
		return item.(V), true
	}

	return *new(V), false
}

// Delete deletes the value for a key.
func (m *SyncMap[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

// Swap swaps the value for a key and returns the previous value if any. The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	var previous1 any
	previous1, loaded = m.Map.Swap(key, value)
	return previous1.(V), loaded
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
func CompareAndDelete[K comparable, V comparable](m *SyncMap[K, V], key K, old V) (deleted bool) {
	return m.Map.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The old value must be of a comparable type.
func CompareAndSwap[K comparable, V comparable](m *SyncMap[K, V], key K, old, new V) (swapped bool) {
	return m.Map.CompareAndSwap(key, old, new)
}
