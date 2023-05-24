package BloomFilter

import (
	"hash"
	"hash/fnv"
	"sync"
)

type ThreadSafeBloomFilter struct {
	bits    []bool
	hashers []hash.Hash64
	mu      sync.RWMutex
}

func NewThreadSafeBloomFilter(m int) *BloomFilter {
	size := optimalSize(m, 0.01)
	k := optimalNumHashers(m, size)
	filter := &BloomFilter{
		bits:    make([]bool, size),
		hashers: make([]hash.Hash64, k),
	}

	// initialise hash functions
	for i := 0; i < k; i++ {
		filter.hashers[i] = fnv.New64()
	}

	return filter
}

func (b *ThreadSafeBloomFilter) Add(data []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, hasher := range b.hashers {
		hasher.Reset()
		hasher.Write(data)
		index := (hasher.Sum64()) % uint64(len(b.bits))
		b.bits[index] = true
	}
}

func (b *ThreadSafeBloomFilter) Contains(data []byte) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, hasher := range b.hashers {
		hasher.Reset()
		// updates hasher with data
		hasher.Write(data)
		index := (hasher.Sum64()) % uint64(len(b.bits))
		if !b.bits[index] {
			return false
		}
	}
	return true
}
