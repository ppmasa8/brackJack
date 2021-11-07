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

func validHitOrStay(s string) bool {
	return s == "hit" || s == "stay"
}

func dealerAction(d *dealer) []int {
	cnt := 0
	cpD := *d
	for _, v := range cpD {
		cnt += v
	}

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
	cntP, cntD := 0, 0
	cpP, cpD := *p, *d
	for _, v := range cpP {
		cntP += v
	}
	for _, v := range cpD {
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

const PWin = 0
const DWin = 1
const Draw = 2

func Judge(p *player, d *dealer) int {
	cntP, cntD := 0, 0
	cpP, cpD := *p, *d
	for _, v := range cpP {
		cntP += v
	}
	for _, v := range cpD {
		cntD += v
	}

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

	fmt.Println(cpP)
	fmt.Println("dealerの手札", cpD)
}
