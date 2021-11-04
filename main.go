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

func validHitOrStay(s string) bool {
	return s == "hit" || s == "stay"
}

func Hit(p []int) {
	c := Rand()
	p = append(p, c)
	for _, v := range p {
		fmt.Println(v)
	}
}

const NoneBurst  = 0
const PBurst     = 1
const DBurst     = 2
const PAndDBurst = 3

func validHands(p, d []int) int {
	cntP, cntD := 0, 0
	for _, v := range p {
		cntP += v
	}
	for _, v := range d {
		cntD += v
	}

	if cntP < 22 && cntD < 22 {
		return NoneBurst
	} else if cntP >= 22 && cntD < 22 {
		return PBurst
	} else if cntP < 22 && cntD >= 22 {
		return DBurst
	} else {
		return PAndDBurst
	}
}

func Judge(p, d []int) bool {

}

func main() {
	var s string

	// プレイヤー
	p := []int{}
	// ディーラー
	d := []int{}

	fmt.Println("あなたの手札")
	for i := 0; i < 2; i++ {
		pn := Rand()
		p = append(p, pn)
		d = append(d, Rand())
		fmt.Println(pn, " ")
	}
	fmt.Println("ヒットするならhit, ステイならstayと入力してください")
	fmt.Scanf("%s", &s)
	if validHitOrStay(s) == false {
		fmt.Println("hitかstayと入力してください", s)
		os.Exit(1)
	}

	if s == "stay" {

	}


}
