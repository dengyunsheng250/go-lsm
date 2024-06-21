package tests

import (
	"lsm"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	b := lsm.NewBloomFilter(50)
	b.Add([]byte("hello"))
	b.Add([]byte("world"))
	t.Log(b.Test([]byte("wor")))
}
