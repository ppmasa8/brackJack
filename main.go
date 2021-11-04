package main

import (
	"fmt"
	"math/rand"
)

func Rand() int {
	n := rand.Intn(12)
	return n + 1
}

func main() {
	a, b := Rand(), Rand()
	fmt.Println("あなたの手札は", a, "と", b, "です。")
}
