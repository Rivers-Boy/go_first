package main

import "testing"

func TestSum(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	got := Sum(arr)
	want := 15

	if got != want {
		t.Errorf("Got '%q', wanted '%q'", got, want)
	}
}
