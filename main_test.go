package main

import (
	"brackjack/lib"
	"fmt"
	"testing"
)

func Test_validHitOrStay(t *testing.T) {
	// True version
	var s string
	s = "stay"
	v := validHitOrStay(s)
	if !v {
		t.Errorf("Return boolean is wrong. %v", v)
	}
	// False version
	var w string
	w = "www"
	x := validHitOrStay(w)
	if x {
		t.Errorf("Return boolean is wrong. %v", x)
	}
}

func Test_dealerAction(t *testing.T) {
	// hit
	d := dealer{1, 2, 3, 4, 5}
	ret := dealerAction(&d)
	if lib.Cnt(ret) <= 15 {
		t.Errorf("Return num is wrong. %v",lib.Cnt(ret))
	}
	// No hit
	nd := dealer{2, 3, 4, 5, 6}
	retd := dealerAction(&nd)
	if lib.Cnt(retd) != 20 {
		t.Errorf("Return num is wrong. %v",lib.Cnt(retd))
	}
}

func Test_validHands(t *testing.T) {
	n, m := player{0, 1}, dealer{10, 11}
	na, ma := validHands(n), validHands(m)
	if na && !ma {
		t.Errorf("Return boolean is wrong. %v", na)
		t.Errorf("Return boolean is wrong. %v", ma)
	}
}

func Test_Judge(t *testing.T) {
	p, d := player{1, 2}, dealer{10, 11}
	ans := Judge(&p, &d)
	if ans != DWin {
		t.Errorf("Return number is wrong. %v", ans)
	}
}

func Test_hitOrStay(t *testing.T) {
	p := player{1, 2}
	ans := hitOrStay(&p)
	fmt.Println(ans)
}
