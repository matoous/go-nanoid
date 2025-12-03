package gonanoid

import (
	"crypto/rand"
	"errors"
	"math"
)

var (
	// defaultAlphabet is the alphabet used for ID characters by default.
	defaultAlphabet = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// AlphaNum is an alphabet of alpha-numerical characters.
	AlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Alpha is an alphabet of upper and lower case letters.
	Alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// AlphaLowerNum is an alphabet of lower case letters and numbers.
	AlphaLowerNum = "abcdefghijklmnopqrstuvwxyz0123456789"

	// AlphaUpperNum is an alphabet of upper case letters and numbers.
	AlphaUpperNum = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// AlphaLower is an alphabet of lower case letters.
	AlphaLower = "abcdefghijklmnopqrstuvwxyz"

	// AlphaUpper is an alphabet of upper case letters.
	AlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Numeric is an alphabet of numerical characters.
	Numeric = "0123456789"

	// CrockfordBase32Upper is the [Crockford](https://www.crockford.com/base32.html) uppercase base32 alphabet.
	CrockfordBase32Upper = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

	// CrockfordBase32Lower is the [Crockford](https://www.crockford.com/base32.html) lower base32 alphabet.
	CrockfordBase32Lower = "0123456789abcdefghjkmnpqrstvwxyz"
)

const (
	defaultSize = 21
)

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
func Generate(alphabet string, size int) (string, error) {
	chars := []rune(alphabet)

	if len(alphabet) == 0 || len(alphabet) > 255 {
		return "", errors.New("alphabet must not be empty and contain no more than 255 chars")
	}
	if size <= 0 {
		return "", errors.New("size must be positive integer")
	}

	mask := getMask(len(chars))
	// estimate how many random bytes we will need for the ID, we might actually need more but this is tradeoff
	// between average case and worst case
	ceilArg := 1.6 * float64(mask*size) / float64(len(alphabet))
	step := int(math.Ceil(ceilArg))

	id := make([]rune, size)
	bytes := make([]byte, step)
	for j := 0; ; {
		_, err := rand.Read(bytes)
		if err != nil {
			return "", err
		}
		for i := 0; i < step; i++ {
			currByte := bytes[i] & byte(mask)
			if currByte < byte(len(chars)) {
				id[j] = chars[currByte]
				j++
				if j == size {
					return string(id[:size]), nil
				}
			}
		}
	}
}

// MustGenerate is the same as Generate but panics on error.
func MustGenerate(alphabet string, size int) string {
	id, err := Generate(alphabet, size)
	if err != nil {
		panic(err)
	}
	return id
}

// New generates secure URL-friendly unique ID.
// Accepts optional parameter - length of the ID to be generated (21 by default).
func New(l ...int) (string, error) {
	var size int
	switch {
	case len(l) == 0:
		size = defaultSize
	case len(l) == 1:
		size = l[0]
		if size < 0 {
			return "", errors.New("negative id length")
		}
	default:
		return "", errors.New("unexpected parameter")
	}
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	id := make([]rune, size)
	for i := 0; i < size; i++ {
		id[i] = []rune(defaultAlphabet)[bytes[i]&63]
	}
	return string(id[:size]), nil
}

// Must is the same as New but panics on error.
func Must(l ...int) string {
	id, err := New(l...)
	if err != nil {
		panic(err)
	}
	return id
}
