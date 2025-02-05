package main

import (
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenPacket("udp", ":8083")
  defer conn.Close()

	if err != nil {
		panic(err)
	}

	for {
    p := make([]byte, 1024)

		_, addr, err := conn.ReadFrom(p)
		if err != nil {
			panic(err)
		}

    text := string(p)
    upper := strings.ToUpper(text)
    upperBytes := []byte(upper)
		_, err = conn.WriteTo(upperBytes, addr)
	}
}
