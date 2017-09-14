package gonanoid

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test the distribution so we are sure, that the collisions won't happen frequently
// Test inspired by AI's javscript nanoid implementation
func TestGenerate(t *testing.T) {
	COUNTER := make(map[byte]int)
	ALPHABET := "abcdefghijklmnopqrstuvwxyz"
	COUNT := 100 * 1000
	SIZE := 5

	Alphabet(ALPHABET)
	Size(SIZE)
	for i := 0; i < COUNT; i++ {
		id := Generate()
		for u := 0; u < len(id); u++ {
			COUNTER[id[u]]++
		}
	}

	for char, count := range COUNTER {
		distribution := (float64(count) * float64(len(ALPHABET))) / float64((COUNT * SIZE))
		if !isInRange(distribution, 0.95, 1.05){
			t.Errorf("distribution error, char %v has %v distribution", char, distribution)
		}
	}
}

// Test if setting the size of nanoid works
func TestSetSize(t *testing.T) {
	assert.NotEqual(t, Size(0), nil, "Function shall return error but returns nil")

	var sizes = []int{4, 10, 20, 22, 30, 40, 60}
	for i := 0; i < len(sizes); i++ {
		Size(sizes[i])
		id := Generate()
		if len(id) != sizes[i] {
			t.Errorf("Nanoid generated with false size: %d, except: %d", len(id), sizes[i])
		}
	}
}

// test if setting the alphabet for nanoid works
func TestAlphabet(t *testing.T) {
	assert.NotEqual(t, Alphabet(""), nil, "Function shall return error but returns nil")
	// test 300 characters
	assert.NotEqual(t, Alphabet("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), nil, "Function shall return error but returns nil")

	var alphabets = []string{"abc", "abcdefg", "abcABC123", "abcdefghABCDEFGH123456_"}
	for _, a := range alphabets {
		CONTAINS := make(map[byte]bool)

		Alphabet(a)
		id := Generate()
		for u := 0; u < len(id); u++ {
			CONTAINS[id[u]] = true
		}

		for key, val := range CONTAINS {
			if val && !byteInString(key, a) {
				t.Errorf("Set alphabet to %v but ID containes latter %v", a, key)
			}
		}
	}
}

// Helping function to find if number is in given range
func isInRange(num float64, from float64, to float64) bool {
	return num > from && num < to
}

// Helping function to find if byte is in given string
func byteInString(b byte, alphabet string) bool{
	for u := 0; u < len(alphabet); u++ {
		if b == alphabet[u]{
			return true
		}
	}
	return false
}


// Benchmark nanoid generator
func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate()
	}
}

// Benchmark generate if alphabet length is not power of 2
func BenchmarkGenerateCustom(b *testing.B) {
	b.StopTimer()
	Alphabet("abcdefghijklmnopqrstuvw")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		Generate()
	}
}