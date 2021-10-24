package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	times := 5
	repeated := Repeat("a", times)
	expected := strings.Repeat("a", times)

	if repeated != expected {
		t.Errorf("got %q, want %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
