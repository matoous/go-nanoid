package gonanoid

import "testing"

var sizes = []int{4, 10, 20, 22, 30, 40, 60}

func TestGenerate(t *testing.T) {
	id := Generate()
	if len(id) != defaultSize {
		t.Errorf("Default Nanoid generated with false size: %d, except: %d", len(id), defaultSize)
	}
	for i := range sizes{
		SetSize(i)
		id = Generate()
		if len(id) != i {
			t.Errorf("Nanoid generated with false size: %d, except: %d", len(id), i)
		}
	}
}

func TestSetSize(t *testing.T) {
	for i := range sizes{
		SetSize(i)
		id := Generate()
		if len(id) != i {
			t.Errorf("Nanoid generated with false size: %d, except: %d", len(id), i)
		}
	}
}

func TestSetAlphabet(t *testing.T) {
	SetAlphabet("abcd")
	if bits != 2 {
		t.Errorf("Set alphabet 'abcd', expect: %d bits, have %d", 2, bits)
	}
	SetAlphabet("abcde")
	if bits != 3 {
		t.Errorf("Set alphabet 'abcde', expect: %d bits, have %d", 3, bits)
	}
	SetAlphabet(defaultAlphabet)
	if bits != 6 {
		t.Errorf("Set alphabet '%s', expect: %d bits, have %d", defaultAlphabet, 6, bits)
	}
}