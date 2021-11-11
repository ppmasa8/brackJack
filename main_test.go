package main

import "testing"

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
