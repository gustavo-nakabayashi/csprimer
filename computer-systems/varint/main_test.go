package main

import (
	"testing"
)

func TestReadUIntFile(t *testing.T) {
	got := ReadUintFile("./1.uint64")
	want := uint64(1)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestEncoding(t *testing.T) {
	tests := []struct {
		filename     string
		expected []byte
	}{
		{"./150.uint64", []byte{0X96, 0X01}},
		{"./1.uint64", []byte{1}},
    {"./maxint.uint64", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01}},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
      want := ReadUintFile(tt.filename)
      got := Decode(Encode(want))
			if got != want {
				t.Errorf("For file %v, got %v, wanted %v",
					tt.filename, got, want)
			}
		})
	}
}
