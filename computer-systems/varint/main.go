package main

import (
	"fmt"
	"math"
)

func AddContinuationBit(bytes []uint8) {
	for i := len(bytes) - 1; i > 0; i-- {
		bytes[i] = bytes[i] + 128
	}
}

func CalcEncodeValue(splited []uint8) int {
	encoded := 0

	for i, v := range splited {
		mult := int(math.Pow(256, float64(i)))
		encoded = encoded + int(v)*mult
	}

	return encoded
}

func Encode(i uint64) int {
	splited := ChunkSplit(i)
	fmt.Print(splited)

	smallEndian := ConvertSmallEndian(splited)

	AddContinuationBit(smallEndian)

	encoded := CalcEncodeValue(smallEndian)

	return encoded
}

func ChunkSplit(i uint64) []uint8 {
	rest := i

	var splited []uint8

	multiplier := uint64(math.Pow(2, 7))

	for {
		if rest == 0 {
			break
		}

		reminder := (rest % multiplier)
		rest = rest / multiplier

		splited = append(splited, uint8(reminder))
	}

	return splited
}

func ConvertSmallEndian(bytes []uint8) []uint8 {
	smallEndian := make([]uint8, len(bytes))

	for i := range smallEndian {
		smallEndian[i] = bytes[len(bytes)-i-1]
	}

	return smallEndian
}
