package gonanoid

import (
	"time"
	"math/rand"
)

const (
	defaultAlphabet = "_~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // URL friendly alphabet
	defaultSize = 22 // default size of Nanoid
	defaultBits = 6
	defaultMask = 1 << defaultBits - 1
	defaultMax = 63/defaultBits
)

// Source of randomness
var (
	src = rand.NewSource(time.Now().UnixNano())
	alphabet = defaultAlphabet
	size = defaultSize
	bits uint64 = defaultBits
	mask int64 = defaultMask
	max uint64 = defaultMax
)

func SetAlphabet(newAlphabet string){
	alphabet = newAlphabet
	bits = computeBits(len(alphabet))
	mask = 1 << bits - 1
	max = 63/bits
}

func SetSize(newSize int){
	size = newSize
}

// Generates new ID 22 characters by default
func Generate() string {
	b := make([]byte, size)
	for i, cache, remain := size-1, src.Int63(), max; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), max
		}
		if idx := int(cache & mask); idx < len(alphabet) {
			b[i] = alphabet[idx]
			i--
		}
		cache >>= bits
		remain--
	}

	return string(b)
}

func computeBits(size int) uint64{
	size--
	var bits uint64 = 0
	for ;size != 0;size >>= 1 {
		bits++;
	}
	return bits
}