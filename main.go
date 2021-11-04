package main

import (
	"fmt"
	"math/rand"
	"os"
)

func Rand() int {
	n := rand.Intn(12)
	return n + 1
}

func Valid(s string) bool {
	return s == "hit" || s == "stay"
}

func main() {
	a, b := Rand(), Rand()
	var s string

	// 対話部分
	fmt.Println("あなたの手札は", a, "と", b, "です。")
	fmt.Println("ヒットするならhit, ステイならstayと入力してください")
	fmt.Scanf("%s", &s)
	if Valid(s) == false {
		fmt.Println("hitかstayと入力してください", s)
		os.Exit(1)
	}

}
