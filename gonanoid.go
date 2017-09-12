package gonanoid

import (
	"crypto/rand"
	"log"
)

const (
	defaultAlphabet = "_~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // URL friendly alphabet
	defaultSize = 22 // default size of Nanoid
	defaultBits = 6 // default bits needed to index default alphabet
	defaultMask = 1 << defaultBits - 1 // default mask for given alphabet
)

var (
	alphabet = defaultAlphabet // alphabet
	size = defaultSize // id size
	bits = defaultBits // bits needed to represent index in alphabet
	mask byte = defaultMask // mask
	bufferSize = int(float64(size)*1.3)
	randomBytes = make([]byte, bufferSize)
)

// Set gonanoid alphabet
func Alphabet(newAlphabet string){
	alphabet = newAlphabet
	bits = computeBits(len(alphabet))
	mask = 1 << uint(bits) - 1
}

// Set generated ids size
func Size(newSize int){
	size = newSize
	bufferSize = int(float64(size)*1.3)
	randomBytes = make([]byte, bufferSize)
}

// Compute bits needed to represent index in array of given size
func computeBits(size int) (bits int){
	for size--; size != 0; size >>= 1 {
		bits++;
	}
	return
}

func Generate() string {
	result := make([]byte, size)
	for i, j := 0, 0; i < size; j++ {
		if j % bufferSize == 0 {
			_, err := rand.Read(randomBytes)
			if err != nil {
				log.Fatal("Unable to generate random bytes")
			}		}
		if idx := int(randomBytes[j % size] & mask); idx < len(alphabet) {
			result[i] = alphabet[idx]
			i++
		}
	}

	return string(result)
}