package gonanoid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasNoCollisions(t *testing.T) {
	tries := 100_000
	used := make(map[string]bool)
	for i := 0; i < tries; i++ {
		id := Must()
		require.False(t, used[id], "shouldn't return colliding IDs")
		used[id] = true
	}
}

func TestFlatDistribution(t *testing.T) {
	tries := 100_000
	alphabet := "abcdefghij"
	size := 10
	chars := make(map[rune]int)
	for i := 0; i < tries; i++ {
		id := MustGenerate(alphabet, size)
		for _, r := range id {
			chars[r]++
		}
	}

	for _, count := range chars {
		require.InEpsilon(t, size*tries/len(alphabet), count, .01, "should have flat distribution")
	}
}

// Benchmark nanoid generator
func BenchmarkNanoid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = New()
	}
}
