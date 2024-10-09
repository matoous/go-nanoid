package gonanoid

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
)

var defaultAlphabet = []rune("_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	defaultSize     = 21
	defaultMaskSize = 5
)

// Generator function
type Generator func([]byte) (int, error)

// BytesGenerator is the default bytes generator
var BytesGenerator Generator = rand.Read

// getMask generates bit mask used to obtain bits from the random bytes that are used to get index of random character
// from the alphabet. Example: if the alphabet has 6 = (110)_2 characters it is sufficient to use mask 7 = (111)_2
func getMask(alphabetSize int) int {
	for i := 1; i <= 8; i++ {
		mask := (2 << uint(i)) - 1
		if mask >= alphabetSize-1 {
			return mask
		}
	}
	return 0
}

// Generate is a low-level function to change alphabet and ID size.
func Generate(rawAlphabet string, size int) (string, error) {
	alphabet := []rune(rawAlphabet)

	if len(alphabet) == 0 || len(alphabet) > 255 {
		return "", fmt.Errorf("alphabet must not empty and contain no more than 255 chars. Current len is %d", len(alphabet))
	}
	if size <= 0 {
		return "", fmt.Errorf("size must be positive integer")
	}

	mask := getMask(len(alphabet))
	ceilArg := 1.6 * float64(mask*size) / float64(len(alphabet))
	step := int(math.Ceil(ceilArg))

	id := make([]rune, size)
	bytes := make([]byte, step)
	for j := 0; ; {
		_, err := BytesGenerator(bytes)
		if err != nil {
			return "", err
		}
		for i := 0; i < step; i++ {
			currByte := bytes[i] & byte(mask)
			if currByte < byte(len(alphabet)) {
				id[j] = alphabet[currByte]
				j++
				if j == size {
					return string(id[:size]), nil
				}
			}
		}
	}
}

// Nanoid generates secure URL-friendly unique ID.
func Nanoid(param ...int) (string, error) {
	var size int
	switch {
	case len(param) == 0:
		size = defaultSize
	case len(param) == 1:
		size = param[0]
		if size < 0 {
			return "", errors.New("negative id length")
		}
	default:
		return "", errors.New("unexpected parameter")
	}
	bytes := make([]byte, size)
	_, err := BytesGenerator(bytes)
	if err != nil {
		return "", err
	}
	id := make([]rune, size)
	for i := 0; i < size; i++ {
		id[i] = defaultAlphabet[bytes[i]&63]
	}
	return string(id[:size]), nil
}

// ID provides more golang idiomatic interface for generating IDs.
// Calling ID is shorter yet still clear `gonanoid.ID(20)` and it requires the lengths parameter by default.
func ID(l int) (string, error) {
	return Nanoid(l)
}

// MustID is the same as ID but panics on error.
func MustID(l int) string {
	id, err := Nanoid(l)
	if err != nil {
		panic(err)
	}
	return id
}

// MustGenerate is the same as Generate but panics on error.
func MustGenerate(rawAlphabet string, size int) string {
	id, err := Generate(rawAlphabet, size)
	if err != nil {
		panic(err)
	}
	return id
}
