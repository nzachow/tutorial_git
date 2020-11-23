package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(40))
}

func sub(x int, y int) int {
	return x - y
}
