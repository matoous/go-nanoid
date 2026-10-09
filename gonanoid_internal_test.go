package gonanoid

import (
	"fmt"
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

func TestNewFlatDistribution(t *testing.T) {
	tries := 100_000
	size := 21
	chars := make(map[byte]int)
	for i := 0; i < tries; i++ {
		id := Must(size)
		for j := 0; j < len(id); j++ {
			chars[id[j]]++
		}
	}

	require.Len(t, chars, len(defaultAlphabet), "all characters in defaultAlphabet should be observed")
	expected := float64(size * tries) / float64(len(defaultAlphabet))
	for _, count := range chars {
		require.InEpsilon(t, expected, float64(count), .03, "should have flat distribution for New")
	}
}

func TestPowerOfTwoDistribution(t *testing.T) {
	testCases := []struct {
		name     string
		alphabet string
		size     int
		tries    int
		epsilon  float64
	}{
		{"binary", "01", 16, 50_000, 0.02},
		{"hex", "0123456789abcdef", 16, 50_000, 0.03},
		{"base32", CrockfordBase32Upper, 16, 50_000, 0.04},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chars := make(map[rune]int)
			for i := 0; i < tc.tries; i++ {
				id := MustGenerate(tc.alphabet, tc.size)
				for _, r := range id {
					chars[r]++
				}
			}

			require.Len(t, chars, len(tc.alphabet))
			expected := float64(tc.size*tc.tries) / float64(len(tc.alphabet))
			for _, count := range chars {
				require.InEpsilon(t, expected, float64(count), tc.epsilon)
			}
		})
	}
}

// Benchmark nanoid generator
func BenchmarkNanoid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = New()
	}
}

func BenchmarkNewSizes(b *testing.B) {
	for _, size := range []int{10, 21, 64} {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, _ = New(size)
			}
		})
	}
}

func BenchmarkGeneratePowerOfTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Generate(defaultAlphabet, 21)
	}
}

func BenchmarkGenerateNonPowerOfTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Generate(AlphaNum, 21)
	}
}
