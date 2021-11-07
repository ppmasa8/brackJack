package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// playerの手札
type player []int

// CPUの手札
type dealer []int

func Rand() int {
	n := rand.New(rand.NewSource(1))
	n.Seed(time.Now().UnixNano())
	i := n.Intn(12)
	return i + 1
}

func Cnt(n []int) int {
	cnt := 0
	for _, v := range n {
		cnt += v
	}
	return cnt
}

func validHitOrStay(s string) bool {
	return s == "hit" || s == "stay"
}

func dealerAction(d *dealer) []int {
	cpD := *d
	cnt := Cnt(cpD)

	if cnt < 16 {
		cpD = append(cpD, Hit())
		cpD = dealerAction(&cpD)
	} else {
		return cpD
	}
	return cpD
}

func Hit() int {
	c := Rand()
	fmt.Println(c)
	return c
}

const NoneBurst = 0
const PBurst = 1
const DBurst = 2
const PAndDBurst = 3

func validHands(p *player, d *dealer) int {
	cpP, cpD := *p, *d
	cntP, cntD := Cnt(cpP), Cnt(cpD)

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

const PWin = 0
const DWin = 1
const Draw = 2

func Judge(p *player, d *dealer) int {
	cpP, cpD := *p, *d
	cntP, cntD := Cnt(cpP), Cnt(cpD)

	if cntP > cntD {
		return PWin
	} else if cntP < cntD {
		return DWin
	} else {
		return Draw
	}
}

func hitOrStay(p *player) []int {
	var s string
	cp := *p
	fmt.Println("ヒットするならhit, ステイならstayと入力してください")
	fmt.Scanf("%s", &s)
	if validHitOrStay(s) == false {
		fmt.Println("hitかstayと入力してください", s)
		os.Exit(1)
	}
	if s == "hit" {
		cp = append(cp, Hit())
		cp = hitOrStay(&cp)
	} else {
		return cp
	}
	return cp
}

func main() {
	cpP, cpD := player{}, dealer{}
	fmt.Println("あなたの手札")
	for i := 0; i < 2; i++ {
		pn := Rand()
		cpP = append(cpP, pn)
		cpD = append(cpD, Rand())
		fmt.Println(pn, " ")
	}

	// hitかstayを選ぶ
	cpP = hitOrStay(&cpP)

	// dealerの行動を決める
	cpD = dealerAction(&cpD)

	// 手札のバースト確認
	v := validHands(&cpP, &cpD)

	// 手札が21を超えた場合
	if v != NoneBurst {
		switch v {
		case PBurst:
			fmt.Println("あなたの手札は21を超えました。")
			fmt.Println("あなたの負け")
			fmt.Println("playerの手札", cpP, Cnt(cpP))
			fmt.Println("dealerの手札", cpD, Cnt(cpD))
			os.Exit(0)
		case DBurst:
			fmt.Println("ディーラーの手札が21を超えました。")
			fmt.Println("あなたの勝ち")
			fmt.Println("playerの手札", cpP, Cnt(cpP))
			fmt.Println("dealerの手札", cpD, Cnt(cpD))
			os.Exit(0)
		case PAndDBurst:
			fmt.Println("あなたの手札は21を超えました。")
			fmt.Println("ディーラーの手札が21を超えました。")
			fmt.Println("引き分けです。")
			fmt.Println("playerの手札", cpP, Cnt(cpP))
			fmt.Println("dealerの手札", cpD, Cnt(cpD))
			os.Exit(0)
		}
	}

	// バーストなしで勝敗を決める場合
	switch Judge(&cpP, &cpD) {
	case PWin:
		fmt.Println("あなたの勝ち")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		os.Exit(0)
	case DWin:
		fmt.Println("あなたの負け")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		os.Exit(0)
	case Draw:
		fmt.Println("引き分けです。")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		os.Exit(0)
	}
}
