package main

import "testing"

func TestHexToRgb(t *testing.T) {
	want := [3]int{ 0, 255, 0 }
	got, _ := HexToRgb("#00ff00")

	if want != got {
		t.Errorf("Wanted: %v, got: %v", want, got)
	}
}
