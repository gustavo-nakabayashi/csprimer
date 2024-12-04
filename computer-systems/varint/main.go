package main

import (
	"encoding/binary"
	"log"
	"os"
)

func ReadUintFile(filename string) uint64 {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Got error reading file: %v", err)
	}

	res := binary.BigEndian.Uint64(content)
	return res
}

func main() {
	ReadUintFile("maxint.uint64")
}

func Encode(i uint64) []byte {
	value := i
	var encodedBytes []byte

	for value > 0 {
		curr := value & 0x7F
		value = value >> 7

		if value > 0 {
			curr = curr + 0x80
		}

		encodedBytes = append(encodedBytes, byte(curr))
	}

	return encodedBytes
}

func Decode(encoded []byte) (decoded uint64) {
	for i := len(encoded) - 1; i >= 0; i-- {
		v := encoded[i]

		if i < len(encoded)-1 {
			decoded = decoded << 7
		}
		x := v & 0x7F
		decoded = decoded | uint64(x)
	}

	return
}
