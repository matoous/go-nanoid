package gonanoid_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TestGenerate(t *testing.T) {
	t.Run("short alphabet", func(t *testing.T) {
		alphabet := ""
		_, err := gonanoid.Generate(alphabet, 32)
		assert.Error(t, err, "should return error if the alphabet is too small")
	})

	t.Run("long alphabet", func(t *testing.T) {
		alphabet := strings.Repeat("a", 256)
		_, err := gonanoid.Generate(alphabet, 32)
		assert.Error(t, err, "should return error if the alphabet is too long")
	})

	t.Run("negative ID length", func(t *testing.T) {
		_, err := gonanoid.Generate("abcdef", -1)
		assert.Error(t, err, "should return error if the requested ID length is invalid")
	})

	t.Run("happy path", func(t *testing.T) {
		alphabet := "abcdef"
		id, err := gonanoid.Generate(alphabet, 6)
		assert.NoError(t, err, "shouldn't return error")
		assert.Len(t, id, 6, "should return ID of requested length")
		for _, r := range id {
			assert.True(t, strings.ContainsRune(alphabet, r), "should use given alphabet")
		}
	})

	t.Run("works with unicode", func(t *testing.T) {
		alphabet := "🚀💩🦄🤖"
		id, err := gonanoid.Generate(alphabet, 6)
		assert.NoError(t, err, "shouldn't return error")
		assert.Equal(t, utf8.RuneCountInString(id), 6, "should return ID of requested length")
		for _, r := range id {
			assert.True(t, strings.ContainsRune(alphabet, r), "should use given alphabet")
		}
	})

	t.Run("single character alphabet", func(t *testing.T) {
		alphabet := "z"
		id, err := gonanoid.Generate(alphabet, 5)
		assert.NoError(t, err)
		assert.Equal(t, "zzzzz", id)
	})

	t.Run("power of two alphabets", func(t *testing.T) {
		testCases := []struct {
			name     string
			alphabet string
			size     int
		}{
			{"size 2", "01", 17},
			{"size 4", "ABCD", 21},
			{"size 8", "01234567", 19},
			{"size 16", "0123456789abcdef", 25},
			{"size 32", gonanoid.CrockfordBase32Upper, 30},
			{"size 64", gonanoid.AlphaNum + "_-", 21},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				id, err := gonanoid.Generate(tc.alphabet, tc.size)
				assert.NoError(t, err)
				assert.Len(t, id, tc.size)
				for _, r := range id {
					assert.True(t, strings.ContainsRune(tc.alphabet, r))
				}
			})
		}
	})

	t.Run("must generate panic", func(t *testing.T) {
		assert.Panics(t, func() {
			gonanoid.MustGenerate("", 10)
		})
	})
}

func TestNew(t *testing.T) {
	t.Run("negative ID length", func(t *testing.T) {
		_, err := gonanoid.New(-1)
		assert.Error(t, err, "should return error if the requested ID length is invalid")
	})

	t.Run("unexpected parameter", func(t *testing.T) {
		_, err := gonanoid.New(10, 20)
		assert.Error(t, err)
	})

	t.Run("zero length", func(t *testing.T) {
		id, err := gonanoid.New(0)
		assert.NoError(t, err)
		assert.Equal(t, "", id)
	})

	t.Run("happy path", func(t *testing.T) {
		id, err := gonanoid.New()
		assert.NoError(t, err, "shouldn't return error")
		assert.Len(t, id, 21, "should return ID of default length")
	})

	t.Run("custom length", func(t *testing.T) {
		id, err := gonanoid.New(6)
		assert.NoError(t, err, "shouldn't return error")
		assert.Len(t, id, 6, "should return ID of requested length")
	})

	t.Run("various lengths", func(t *testing.T) {
		lengths := []int{1, 2, 3, 4, 5, 6, 7, 8, 15, 16, 21, 32, 64, 100}
		for _, length := range lengths {
			id, err := gonanoid.New(length)
			assert.NoError(t, err)
			assert.Len(t, id, length)
		}
	})

	t.Run("must panic on invalid length", func(t *testing.T) {
		assert.Panics(t, func() {
			gonanoid.Must(-1)
		})
	})
}
