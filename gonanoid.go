package gonanoid

import (
	"crypto/rand"
	"math"
)

const (
	defaultAlphabet = "_~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // URL friendly alphabet
	defaultSize     = 22                                                                 // default size of Nanoid
	defaultBits     = 6                                                                  // default bits needed to index default alphabet
	defaultMask     = 1<<defaultBits - 1                                                 // default mask for given alphabet
)

var (
	alphabet      = defaultAlphabet // alphabet
	size          = defaultSize     // id size
	mask     byte = defaultMask     // mask
	step          = int(math.Ceil(1.6 * float64(mask) * 22.0 / 64.0))
	buffer        = make([]byte, step)
)

// Set nanoid alphabet
func Alphabet(newAlphabet string) {
	alphabet = newAlphabet
	mask = 1<<computeBits(len(alphabet)) - 1
	step = int(math.Ceil(1.6 * float64(mask) * float64(size) / float64(len(alphabet))))
	buffer = make([]byte, step)
}

// Set nanoid size
func Size(newSize int) {
	size = newSize
	step = int(math.Ceil(1.6 * float64(mask) * float64(size) / float64(len(alphabet))))
	buffer = make([]byte, step)
}

// Compute bits needed to represent index in array of given size
func computeBits(size int) (bits uint) {
	for size--; size != 0; size >>= 1 {
		bits++;
	}
	return
}

// Generate nanoid
func Generate() string {
	result := make([]byte, size)
	for u := 0; ; {
		if _, err := rand.Read(buffer); err != nil {
			panic(err)
		}
		for i := 0; i < step; i++ {
			b := buffer[i] & mask
			if b < byte(len(alphabet)) {
				result[u] = alphabet[b]
				u++
				if u == size {
					return string(result)
				}
			}
		}
	}
}
