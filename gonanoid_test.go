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
}

func TestNew(t *testing.T) {
	t.Run("negative ID length", func(t *testing.T) {
		_, err := gonanoid.New(-1)
		assert.Error(t, err, "should return error if the requested ID length is invalid")
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
}

func TestGenerateAlphabetCharacterLimit(t *testing.T) {
	for _, tc := range []struct {
		name      string
		start     rune
		count     int
		wantError bool
	}{
		{"empty", 'a', 0, true},
		{"single Unicode character", '🚀', 1, false},
		{"63 Unicode characters", '😀', 63, false},
		{"64 Unicode characters", '😀', 64, false},
		{"255 Unicode characters", '😀', 255, false},
		{"256 Unicode characters", '😀', 256, true},
		{"127 ASCII characters", 1, 127, false},
		{"256 characters", 1, 256, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			chars := make([]rune, tc.count)
			for i := range chars {
				chars[i] = tc.start + rune(i)
			}
			alphabet := string(chars)
			id, err := gonanoid.Generate(alphabet, 32)
			if tc.wantError {
				assert.Error(t, err)
				return
			}
			if !assert.NoError(t, err, "%d characters in %d bytes", len(chars), len(alphabet)) {
				return
			}
			assert.Equal(t, 32, utf8.RuneCountInString(id))
			for _, r := range id {
				assert.True(t, strings.ContainsRune(alphabet, r))
			}
		})
	}
}
