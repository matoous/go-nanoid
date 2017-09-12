package main

import (
	"fmt"
	"go-nanoid"
)

func main(){
	id := gonanoid.Generate()
	fmt.Printf("Generated default id: %s\n", id)
}