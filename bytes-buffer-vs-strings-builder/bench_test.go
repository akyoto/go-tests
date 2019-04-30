package stringconcat_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	_1 = "unknown string"
	_2 = "another string"
	_3 = "yet another string"
	_4 = "oh good lord yet another string"
)

func BenchmarkBuilderWriteString(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		var builder strings.Builder
		builder.WriteString(_1)
		builder.WriteString(_2)
		builder.WriteString(_3)
		builder.WriteString(_4)
		_ = builder.String()
	}
}

func BenchmarkBuilderFprintf(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		var builder strings.Builder
		fmt.Fprintf(&builder, "%s%s%s%s", _1, _2, _3, _4)
		_ = builder.String()
	}
}

func BenchmarkBufferWriteString(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		buffer.WriteString(_1)
		buffer.WriteString(_2)
		buffer.WriteString(_3)
		buffer.WriteString(_4)
		_ = buffer.String()
	}
}

func BenchmarkBufferFprintf(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		fmt.Fprintf(&buffer, "%s%s%s%s", _1, _2, _3, _4)
		_ = buffer.String()
	}
}
