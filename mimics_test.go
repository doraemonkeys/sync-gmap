package syncgmap

import (
	"testing"
	"unsafe"
)

func TestNewSyncMap(t *testing.T) {
	m1 := &SyncMap[string, int]{}
	m1Addr := &m1.Map
	if unsafe.Pointer(m1) != unsafe.Pointer(m1Addr) {
		t.Errorf("SyncMap Address = %v, want %v", unsafe.Pointer(m1), unsafe.Pointer(m1Addr))
	}
}

func TestSyncMap_Iterate(t *testing.T) {
	m1 := &SyncMap[string, int]{}
	m1.Store("key1", 1)
	m1.Store("key2", 2)
	m1.Store("key3", 3)

	for k, v := range m1.Range {
		t.Logf("key: %v, value: %v", k, v)
	}
}
