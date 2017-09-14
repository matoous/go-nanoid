package gonanoid

import (
	"crypto/rand"
	"math"
	"github.com/go-errors/errors"
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

// Alphabet - sets nanoid alphabet to given string
// if alphabet length is power of two,
// simpler and faster algorithm will be used,
// because we don't need to skip any masked bytes
func Alphabet(newAlphabet string) error {
	length := len(newAlphabet)
	if length < 1 || length > 256 {
		return errors.New("alphabet length must be between 1 and 256")
	}
	// set values
	alphabet = newAlphabet
	mask = 1<<computeBits(length) - 1
	recompute()
	return nil
}

// Size - sets size of generated nanoid
func Size(newSize int) error {
	if newSize < 1 {
		return errors.New("size must be greater than 0")
	}
	size = newSize
	recompute()
	return nil
}

// Generate - generates nanoid
func Generate() (string, error) {
	result := make([]byte, size)
	// speed up when it is possible
	if simple {
		buff := make([]byte, size)
		if _, err := rand.Read(buff); err != nil {
			return "", err
		}
		for i := 0; i < size; i++ {
			result[i] = alphabet[buff[i]&mask]
		}
		return string(result), nil
	} else {
		buff := make([]byte, step)
		for u := 0; ; {
			if _, err := rand.Read(buff); err != nil {
				return "", err
			}
			for i := 0; i < step; i++ {
				b := buff[i] & mask
				if b < byte(len(alphabet)) {
					result[u] = alphabet[b]
					u++
					if u == size {
						return string(result), nil
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

// recompute random bytes array length if needed
func recompute(){
	// check simple algo availability
	if isPowerOfTwo(len(alphabet)) {
		simple = true
	} else {
		simple = false
		step = computeStep(mask, size, len(alphabet))
	}
}

// compute step
func computeStep(mask byte, size, length int) int {
	return int(math.Ceil(1.6 * float64(mask) * float64(size) / float64(length)))
}