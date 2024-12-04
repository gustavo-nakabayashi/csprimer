package main

import (
	"math"
	"reflect"
	"testing"
)

func TestChunkSplit(t *testing.T) {
	got := ChunkSplit(150)
	want := []uint8{22, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestConvertToSmallEndian(t *testing.T) {
	got := ConvertSmallEndian([]uint8{22,1})
	want := []uint8{1, 22}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestAddContinuationBit(t *testing.T) {
	got := []uint8{1, 22}
	AddContinuationBit(got)
	want := []uint8{1, 150}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCalcEncodeValue(t *testing.T) {
	got := CalcEncodeValue([]uint8{1, 150})
	want := 0x9601

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestEncoding(t *testing.T) {
    tests := []struct {
        name     string
        input    uint64
        expected int
    }{
        {"encode_150", 150, 0x9601},
        {"encode_1", 1, 0x01},
        {"encode_127", 127, 0x7F},         // Max 7-bit number
        {"encode_128", 128, 0x8001},       // Min 2-byte number
        {"encode_16383", 16383, 0xFF7F},   // Max 14-bit number
        {"encode_16384", 16384, 0x808001}, // Min 3-byte number
        {"encode_300", 300, 0xAC02},       // Random 2-byte number
        {"encode_0", 0, 0x00},             // Edge case: zero
        {"encode_255", 255, 0xFF01},       // One byte boundary
        {"encode_256", 256, 0x8002},       // Two byte boundary
        {"encode_2097151", 2097151, 0xFFFF7F}, // Max 21-bit number
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Encode(tt.input)
            if got != tt.expected {
                t.Errorf("Encode(%d) = 0x%X, want 0x%X", 
                    tt.input, got, tt.expected)
            }
        })
    }
}

