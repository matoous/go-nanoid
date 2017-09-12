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
	defaultMax = 63/defaultBits // default number of values available from 63 random bits
)

var (
	alphabet = defaultAlphabet // alphabet
	size = defaultSize // id size
	bits uint64 = defaultBits // bits needed to represent index in alphabet
	mask byte = defaultMask // mask
	max uint64 = defaultMax // available values from 63 random bits
)

// Set gonanoid alphabet
func SetAlphabet(newAlphabet string){
	alphabet = newAlphabet
	bits = computeBits(len(alphabet))
	mask = 1 << bits - 1
	max = 63/bits
}

// Set generated ids size
func SetSize(newSize int){
	size = newSize
}

// compute bits needed to represent index in array of given size
func computeBits(size int) uint64{
	size--
	var bits uint64 = 0
	for ;size != 0;size >>= 1 {
		bits++;
	}
	return bits
}

func Generate() string {
	result := make([]byte, size)
	bufferSize := int(float64(size)*1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < size; j++ {
		if j%bufferSize == 0 {
			randomBytes = secureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%size] & mask); idx < len(alphabet) {
			result[i] = alphabet[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func secureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}