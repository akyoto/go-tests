package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/OneOfOne/xxhash"
)

const (
	collisionIterations = 25
	dataSizeMin         = 8
	dataSizeMax         = 10000
)

func TestNoCollisions(t *testing.T) {
	hashes := map[uint64]bool{}

	// Generate payloads
	for dataSize := dataSizeMin; dataSize < dataSizeMax; dataSize++ {
		for i := 0; i < collisionIterations; i++ {
			data := make([]byte, dataSize)
			_, _ = rand.Read(data)
			h := Bytes(data)
			_, exists := hashes[h]

			if exists {
				t.Fatalf("Hash '%d' for '%s' already exists", h, data)
			}

			hashes[h] = true
			fmt.Println(h)
		}
	}
}

func TestNoCollisionsXXHash(t *testing.T) {
	hashes := map[uint64]bool{}

	// Generate payloads
	for dataSize := dataSizeMin; dataSize < dataSizeMax; dataSize++ {
		for i := 0; i < collisionIterations; i++ {
			data := make([]byte, dataSize)
			_, _ = rand.Read(data)
			hasher := xxhash.New64()
			_, _ = hasher.Write(data)
			h := hasher.Sum64()
			_, exists := hashes[h]

			if exists {
				t.Fatalf("Hash '%d' for '%s' already exists", h, data)
			}

			hashes[h] = true
		}
	}
}

func BenchmarkBytes(b *testing.B) {
	data := bytes.Repeat([]byte("Hello World"), 1000)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Bytes(data)
		}
	})
}

func BenchmarkXXHash(b *testing.B) {
	data := bytes.Repeat([]byte("Hello World"), 1000)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hasher := xxhash.New64()
			_, _ = hasher.Write(data)
			hasher.Sum64()
		}
	})
}
