package leetcode

import "testing"

func TestMaximumSwap(t *testing.T) {
	test1 := 1234
	want := 4231
	got := maximumSwap(test1)
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestMaximumSwap2(t *testing.T) {
	test1 := 841882
	want := 881842
	got := maximumSwap(test1)
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}
