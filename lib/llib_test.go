package lib

import "testing"

func Test_rand(t *testing.T) {
	n := Rand()
	if n > 13 || n < 0 {
		t.Errorf("Return num is wrong. result=%v", n)
	}
}

func Test_cnt(t *testing.T) {
	n := []int{1, 2, 3, 4, 5}
	cnt := Cnt(n)
	if cnt != 15 {
		t.Errorf("Return cnt is wrong. result=%v", cnt)
	}
}
