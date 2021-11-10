package main

import (
	"math/rand"
	"time"
)

func Rand() int {
	n := rand.New(rand.NewSource(1))
	n.Seed(time.Now().UnixNano())
	return n.Intn(12) + 1
}

func Cnt(n []int) (cnt int) {
	for _, v := range n {
		cnt += v
	}
	return cnt
}
