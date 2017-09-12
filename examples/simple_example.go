package main

import (
	"gonanoid"
	"fmt"
)

func main(){
	id := gonanoid.Generate()
	fmt.Printf("Generated default id: %s\n", id)
}