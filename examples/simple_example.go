package main

import (
	"fmt"

	"github.com/matoous/go-nanoid"
)

func main() {
	// Simple usage
	id, err := gonanoid.Nanoid()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)

	// Custom length
	id, err = gonanoid.ID(5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)

	// Custom alphabet
	id, err = gonanoid.Generate("abcdefg", 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)

	// Custom non ascii alphabet
	id, err = gonanoid.Generate("こちんにабдежиклмнは你好喂שלום😯😪🥱😌😛äöüß", 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)
}
