package syncmap

import (
	"sync"
)

const (
	//DefaultSyncCaps indicate num of syncMapSlice in syncMap
	DefaultSyncCaps uint32 = 32
	fnvOffsetBasis  uint32 = 2166136261
	fncPrime        uint32 = 16777619
)

type syncMapSlice struct {
	mu          sync.Mutex
	schema2conn map[string]int
}

func newSyncMapSlice() *syncMapSlice {
	return &syncMapSlice{schema2conn: make(map[string]int)}
}

func (o *syncMapSlice) getAndAdd(key string, value int) int {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.schema2conn[key] += value
	return o.schema2conn[key]
}

// SyncMap act as a concurrent map.
type SyncMap struct {
	mapSlice []*syncMapSlice
	cap      uint32
}

// NewSyncMap returns a concurrent map.
func NewSyncMap(capacity uint32) *SyncMap {
	sm := &SyncMap{
		mapSlice: make([]*syncMapSlice, capacity),
		cap:      capacity,
	}
	for i := 0; i < int(capacity); i++ {
		sm.mapSlice[i] = newSyncMapSlice()
	}
	return sm
}

// GetAndAdd add map[key] with value
func (o *SyncMap) GetAndAdd(key string, value int) int {
	hashVal := fnv32(key)
	return o.mapSlice[hashVal%o.cap].getAndAdd(key, value)
}

// FNV1 hash
func fnv32(key string) uint32 {
	hash := uint32(fnvOffsetBasis)
	const prime32 = uint32(fncPrime)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
