package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		var n uint8 = b[0]

		if n > 48 && n < 58 {
			beepTimes := int(n) - 48
			for i := 0; i < beepTimes; i++ {
				os.Stdout.Write([]byte{7})
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}
