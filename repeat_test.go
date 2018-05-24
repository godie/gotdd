package main

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expexted := "aaaaaaaaaa"

	if repeated != expexted {
		t.Errorf("expected '%s' but got '%s'", expexted, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
