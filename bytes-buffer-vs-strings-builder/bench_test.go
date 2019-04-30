package stringconcat_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkA(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		var builder strings.Builder
		fmt.Fprintf(&builder, "%d...%s", 13, "Hello")
		builder.WriteString("EOL")
		_ = builder.String()
	}
}

func BenchmarkB(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		fmt.Fprintf(&buffer, "%d...%s", 13, "Hello")
		buffer.WriteString("EOL")
		_ = buffer.String()
	}
}
