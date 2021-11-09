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
	return n.Intn(12) + 1
}

func Cnt(n []int) (cnt int) {
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
	}

	return cpD
}

func Hit() int {
	return Rand()
}

func validHands(n []int) bool {
	cnt := Cnt(n)
	return cnt < 22
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
		return []int{}
	}
	if s == "hit" {
		cp = append(cp, Hit())
		fmt.Println(cp, "計", Cnt(cp))
		if validHands(cp) {
			cp = hitOrStay(&cp)
		} else {
			return cp
		}
	}
	return cp
}

func process() {
	cpP, cpD := player{}, dealer{}
	fmt.Println("あなたの手札")
	for i := 0; i < 2; i++ {
		pn := Rand()
		cpP = append(cpP, pn)
		cpD = append(cpD, Rand())
	}
	fmt.Println(cpP, "計", Cnt(cpP))

	// hitかstayを選ぶ
	cpP = hitOrStay(&cpP)

	// hit or stay 以外の文字列入力を受けた際のハンドリング
	if len(cpP) == 0 {
		return
	}

	// dealerの行動を決める
	cpD = dealerAction(&cpD)

	// 手札のバースト確認
	bp, bd := validHands(cpP), validHands(cpD)

	// 手札が21を超えた場合
	if !bp && !bd {
		fmt.Println("あなたの手札は21を超えました。")
		fmt.Println("ディーラーの手札が21を超えました。")
		fmt.Println("引き分けです。")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		return
	} else if !bp && bd {
		fmt.Println("あなたの手札は21を超えました。")
		fmt.Println("あなたの負け")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		return
	} else if bp && !bd {
		fmt.Println("ディーラーの手札が21を超えました。")
		fmt.Println("あなたの勝ち")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		return
	}

	// バーストなしで勝敗を決める場合
	switch Judge(&cpP, &cpD) {
	case PWin:
		fmt.Println("あなたの勝ち")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		return
	case DWin:
		fmt.Println("あなたの負け")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		return
	case Draw:
		fmt.Println("引き分けです。")
		fmt.Println("playerの手札", cpP, Cnt(cpP))
		fmt.Println("dealerの手札", cpD, Cnt(cpD))
		return
	}
}

func retry() {
	var s string
	fmt.Println("retryしますか？yes/no")
	fmt.Scanf("%s", &s)
	if s == "yes" {
		process()
		retry()
	}
}

func main() {
	process()
	retry()
	os.Exit(0)
}
