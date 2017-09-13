package main

import (
	"fmt"
	"github.com/matoous/go-nanoid"
)

func main(){
	id := gonanoid.Generate()
	fmt.Printf("Generated default id: %s\n", id)
}
