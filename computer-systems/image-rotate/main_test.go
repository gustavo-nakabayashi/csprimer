package main

import "testing"

func TestReadBmpData(t *testing.T) {
	got := ReadBmpData("./stretch-goal.bmp")
	want := BmpImage{width: 420, height: 420, bitsPerPixel: 24, startingAddress: 138, bytesPerRow: 1260}

	if got != want {
		t.Errorf("wanted: %d, got : %d", want, got)
	}
}
