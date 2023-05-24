package BloomFilter

import (
	"hash"
	"hash/fnv"
)

type BloomFilter struct {
	bits    []bool
	hashers []hash.Hash64
}

func NewBloomFilter(m int) *BloomFilter {
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

func (b *BloomFilter) Add(data []byte) {
	for _, hasher := range b.hashers {
		hasher.Reset()
		hasher.Write(data)
		index := hasher.Sum64() % uint64(len(b.bits))
		b.bits[index] = true
	}
}

func (b *BloomFilter) Contains(data []byte) bool {
	for _, hasher := range b.hashers {
		hasher.Reset()
		hasher.Write(data)
		index := hasher.Sum64() % uint64(len(b.bits))
		if !b.bits[index] {
			return false
		}
	}
	return true
}
