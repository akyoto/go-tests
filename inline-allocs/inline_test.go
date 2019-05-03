package allocs_test

import (
	"bytes"
	"testing"
)

func BenchmarkInline(b *testing.B) {
	var buffer bytes.Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buffer.Write([]byte("Hello World"))
	}
}

func BenchmarkOptimized(b *testing.B) {
	var buffer bytes.Buffer
	newline := []byte("Hello World")
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buffer.Write(newline)
	}
}
