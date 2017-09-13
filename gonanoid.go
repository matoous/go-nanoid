package gonanoid

import (
	"crypto/rand"
	"math"
)

const (
	defaultAlphabet = "_~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // URL friendly alphabet
	defaultSize     = 22                                                                 // default size of Nanoid
)

var (
	alphabet      = defaultAlphabet                   // alphabet
	size          = defaultSize                       // id size
	mask     byte = 1<<computeBits(len(alphabet)) - 1 // mask
	simple        = true
	step     int
)

// Alphabet sets nanoid alphabet to given string
// if alphabet length is power of two,
// simpler and faster algorithm will be used,
// because we don't need to skip any masked bytes
func Alphabet(newAlphabet string) {
	alphabet = newAlphabet
	mask = 1<<computeBits(len(alphabet)) - 1
	// check simple algo availability
	if isPowerOfTwo(len(alphabet)) {
		simple = true
	} else {
		simple = false
		step = int(math.Ceil(1.6 * float64(mask) * float64(size) / float64(len(alphabet))))
	}
}

// Size sets size of generated nanoid
func Size(newSize int) {
	size = newSize
}

// Generate nanoid
func Generate() string {
	result := make([]byte, size)
	// speed up when it is possible
	if simple {
		buff := make([]byte, size)
		if _, err := rand.Read(buff); err != nil {
			panic(err)
		}
		for i := 0; i < size; i++ {
			result[i] = alphabet[buff[i]&mask]
		}
		return string(result)
	} else {
		buff := make([]byte, step)
		for u := 0; ; {
			if _, err := rand.Read(buff); err != nil {
				panic(err)
			}
			for i := 0; i < step; i++ {
				b := buff[i] & mask
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
}


// helpers

// find out if number n is power of 2
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for n != 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

// compute bits needed to represent index in alphabet of given size
func computeBits(size int) (bits uint) {
	for size--; size != 0; size >>= 1 {
		bits++;
	}
	return
}
