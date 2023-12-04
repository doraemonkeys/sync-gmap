package syncgmap

func (m *SyncMap[K, V]) Len() int {
	len := 0
	m.Map.Range(func(key, value any) bool {
		len++
		return true
	})

	return len
}

func (m *SyncMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	m.Range(func(key K, value V) bool {
		keys = append(keys, key)
		return true
	})

	return keys
}

func (m *SyncMap[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	m.Range(func(key K, value V) bool {
		values = append(values, value)
		return true
	})

	return values
}

func (m *SyncMap[K, V]) Clear() {
	m.Map.Range(func(key, value any) bool {
		m.Map.Delete(key)
		return true
	})
}

func (m *SyncMap[K, V]) Clone() *SyncMap[K, V] {
	clone := new(SyncMap[K, V])
	m.Range(func(key K, value V) bool {
		clone.Store(key, value)
		return true
	})

	return clone
}

func (m *SyncMap[K, V]) Merge(other *SyncMap[K, V]) {
	if other == nil {
		return
	}
	other.Range(func(key K, value V) bool {
		m.Store(key, value)
		return true
	})
}
