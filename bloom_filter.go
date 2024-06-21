package lsm

import "github.com/spaolacci/murmur3"

func NewBloomFilter(size int64) *BloomFilter {
	return &BloomFilter{
		Size:   size,
		Bitset: make([]bool, size),
	}
}

func (b *BloomFilter) Add(item []byte) {
	hash1 := murmur3.Sum64(item)
	hash2 := murmur3.Sum64WithSeed(item, 1)
	hash3 := murmur3.Sum64WithSeed(item, 2)
	b.Bitset[hash1%uint64(b.Size)] = true
	b.Bitset[hash2%uint64(b.Size)] = true
	b.Bitset[hash3%uint64(b.Size)] = true
}

func (b *BloomFilter) Test(item []byte) bool {
	hash1 := murmur3.Sum64(item)
	hash2 := murmur3.Sum64WithSeed(item, 1)
	hash3 := murmur3.Sum64WithSeed(item, 2)
	return b.Bitset[hash1%uint64(b.Size)] &&
		b.Bitset[hash2%uint64(b.Size)] &&
		b.Bitset[hash3%uint64(b.Size)]
}
