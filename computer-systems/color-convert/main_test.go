package main

import "testing"

func TestHexToRgb(t *testing.T) {
	testTable := map[string]string{
		"#00ff00":   "rgb(0 255 0)",
		"#123":      "rgb(17 34 51)",
		"#00FF00":   "rgb(0 255 0)",
		"#0000FFc0": "rgba(0 0 255 / 0.75294)",
		"#00f8":     "rgba(0 0 255 / 0.53333)",
	}

	for testValue, want := range testTable {
		t.Run(testValue, func(t *testing.T) {
			got, _ := HexToRgb(testValue)

			if want != got {
				t.Errorf("Wanted: %v, got: %v", want, got)
			}
		})
	}
}

