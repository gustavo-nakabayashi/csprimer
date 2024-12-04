/*

 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
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

func HexToRgb(hex string) ([3]int, error) {
	if len(hex) != 7 {
		return [3]int{-1, -1, -1}, errors.New("Hex is not in the format #XXXXXX")
	}

	var rgb [3]int

	for i := 0; i < 3; i++ {
		hexColor := hex[1+i*2 : 3+i*2]
		color0, ok := hexCharToInt[hexColor[0:1]]
		if !ok {
			log.Fatalf("Failed to convert to number")
		}

		color1, ok := hexCharToInt[hexColor[1:2]]
		if !ok {
			log.Fatalf("Failed to convert to number")
		}

		rgb[i] = color0*16 + color1
	}

	return rgb, nil
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

		rgbString := fmt.Sprintf("rgb(%d, %d, %d)", rgb[0], rgb[1], rgb[2])

		converted := hexRegex.ReplaceAllLiteralString(line, rgbString)
		os.Stdout.WriteString(converted)
		os.Stdout.WriteString("\n")
	}
}
