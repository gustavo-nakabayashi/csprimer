/*
Normalize 3 and 4 char hexes
implement alpha
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

var hexCharToInt = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"a": 10,
	"b": 11,
	"c": 12,
	"d": 13,
	"e": 14,
	"f": 15,
}

func ConvertColorValues() {}

func NormalizeHex(hex string) string {
	normalized := strings.ToLower(hex)

	if len(hex) == 7 || len(hex) == 9 {
		return normalized
	}

	normalized = normalized[:1] + strings.Repeat(normalized[1:2], 2) + strings.Repeat(normalized[2:3], 2) + strings.Repeat(normalized[3:4], 2) + strings.Repeat(normalized[4:], 2)
	return normalized
}

func CalcRgb(hex string) []int {
	digits := hex[1:]

	rgb := make([]int, len(digits)/2)

	for i := 0; i < len(digits); i = i + 2 {
		hexColor := digits[i : i+2]
		color0, ok := hexCharToInt[hexColor[0:1]]
		if !ok {
			log.Fatalf("Failed to convert to number")
		}

		color1, ok := hexCharToInt[hexColor[1:2]]
		if !ok {
			log.Fatalf("Failed to convert to number")
		}

		rgb[i/2] = color0<<4 + color1
	}

	return rgb
}

func HexToRgb(hex string) (string, error) {
	if !slices.Contains([]int{4, 5, 7, 9}, len(hex)) {
		return "", errors.New("Hex is not valid")
	}

	normalized := NormalizeHex(hex)

	rgb := CalcRgb(normalized)

	var rgbString string

	if len(rgb) == 3 {
		rgbString = fmt.Sprintf("rgb(%v %v %v)", rgb[0], rgb[1], rgb[2])
	} else {
		rgbString = fmt.Sprintf("rgba(%v %v %v / %.5f)", rgb[0], rgb[1], rgb[2], float32(rgb[3])/255.0)
	}

	return rgbString, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hexRegex := regexp.MustCompile(`#[0-9a-fA-F]+`)

	for scanner.Scan() {
		line := scanner.Text()

		hex := hexRegex.FindString(line)

		if hex == "" {
			os.Stdout.WriteString(line)
			os.Stdout.WriteString("\n")
			continue
		}

		rgb, err := HexToRgb(hex)
		if err != nil {
			log.Fatal(err)
		}

		converted := hexRegex.ReplaceAllLiteralString(line, rgb)
		os.Stdout.WriteString(converted)
		os.Stdout.WriteString("\n")
	}
}
