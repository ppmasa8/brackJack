package main

import (
	"brackjack/lib"
	"fmt"
	"os"
)

// playerの手札
type player []int

// CPUの手札
type dealer []int

func validHitOrStay(s string) bool {
	return s == "hit" || s == "stay"
}

func dealerAction(d *dealer) []int {
	cpD := *d
	cnt := lib.Cnt(cpD)

	if cnt < 16 {
		cpD = append(cpD, Hit())
		cpD = dealerAction(&cpD)
	}

	return cpD
}

func Hit() int {
	return lib.Rand()
}

func validHands(n []int) bool {
	cnt := lib.Cnt(n)
	return cnt > 1 && cnt < 22
}

const PWin = 0
const DWin = 1
const Draw = 2

func Judge(p *player, d *dealer) int {
	cpP, cpD := *p, *d
	cntP, cntD := lib.Cnt(cpP), lib.Cnt(cpD)

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
	fmt.Println(lib.ChooseHitOrStay)
	fmt.Scanf("%s", &s)
	if validHitOrStay(s) == false {
		fmt.Println(lib.InputHitOrStay, s)
		return []int{}
	}
	if s == "hit" {
		cp = append(cp, Hit())
		fmt.Println(cp, "計", lib.Cnt(cp))
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
	fmt.Println(lib.PlayerHands)
	for i := 0; i < 2; i++ {
		pn := lib.Rand()
		cpP = append(cpP, pn)
		cpD = append(cpD, lib.Rand())
	}
	fmt.Println(cpP, "計", lib.Cnt(cpP))

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
		fmt.Println(lib.BurstPlayer)
		fmt.Println(lib.BurstDealer)
		fmt.Println(lib.DRAW)
		fmt.Println(lib.PlayerHands, cpP, lib.Cnt(cpP))
		fmt.Println(lib.DealerHands, cpD, lib.Cnt(cpD))
		return
	} else if !bp && bd {
		fmt.Println(lib.BurstPlayer)
		fmt.Println(lib.PlayerLose)
		fmt.Println(lib.PlayerHands, cpP, lib.Cnt(cpP))
		fmt.Println(lib.DealerHands, cpD, lib.Cnt(cpD))
		return
	} else if bp && !bd {
		fmt.Println(lib.BurstDealer)
		fmt.Println(lib.PlayerWin)
		fmt.Println(lib.PlayerHands, cpP, lib.Cnt(cpP))
		fmt.Println(lib.DealerHands, cpD, lib.Cnt(cpD))
		return
	}

	// バーストなしで勝敗を決める場合
	switch Judge(&cpP, &cpD) {
	case PWin:
		fmt.Println(lib.PlayerWin)
		fmt.Println(lib.PlayerHands, cpP, lib.Cnt(cpP))
		fmt.Println(lib.DealerHands, cpD, lib.Cnt(cpD))
		return
	case DWin:
		fmt.Println(lib.PlayerLose)
		fmt.Println(lib.PlayerHands, cpP, lib.Cnt(cpP))
		fmt.Println(lib.DealerHands, cpD, lib.Cnt(cpD))
		return
	case Draw:
		fmt.Println(lib.DRAW)
		fmt.Println(lib.PlayerHands, cpP, lib.Cnt(cpP))
		fmt.Println(lib.DealerHands, cpD, lib.Cnt(cpD))
		return
	}
}

func retry() {
	var s string
	fmt.Println(lib.ChooseRetry)
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
