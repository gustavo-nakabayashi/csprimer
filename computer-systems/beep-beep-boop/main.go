package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// scanner := bufio.NewReader(os.Stdin)
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	// exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		var n uint8 = b[0]
		switch n {
		case 27, 113, 81:
			fmt.Println("ESC, or 'q' or 'Q' was hitted!")
			break
		default:
			fmt.Printf("You typed : %d\n", n)
		}
	}

	// for {
	// 	content, _, err := scanner.ReadRune()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	if content > 48 && content < 58 {
	// 		beepTimes := int(content) - 48
	// 		for i := 0; i < beepTimes; i++ {
	// 			os.Stdout.Write([]byte{7})
	// 			time.Sleep(200 * time.Millisecond)
	// 		}
	// 	}
	// }
}
