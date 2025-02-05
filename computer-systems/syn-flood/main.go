package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type PcapHeader struct {
	MagicNumber         uint32
	MajorVersion        uint16
	MinorVersion        uint16
	TimeZoneOffset      uint32
	TimestampAccuracy   uint32
	SnapshotLenght      uint32
	LinkLayerHeaderType uint32
}

type PcapPacketHeader struct {
	Timestamp     uint32
	TimestampMs   uint32
	Length        uint32
	UntruncLength uint32
}

type ProtocolType struct {
	ProtocolType uint32
}

type IpHeader struct {
	_      uint16
	Length uint16
	_      uint64
	_      uint64
}

func main() {
	file, err := os.ReadFile("./synflood.pcap")
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(file)

	var pcapHeader PcapHeader

	err = binary.Read(r, binary.LittleEndian, &pcapHeader)
	if err != nil {
		panic(err)
	}

	if pcapHeader.MagicNumber != 0xa1b2c3d4 || pcapHeader.MajorVersion != 2 || pcapHeader.MinorVersion != 4 || pcapHeader.LinkLayerHeaderType != 0 {
		panic("Error parsing header")
	}

	packetsCount := 0
	synCount := 0.0
	ackCount := 0.0

	for {
		var pcapPacketHeader PcapPacketHeader
		err := binary.Read(r, binary.LittleEndian, &pcapPacketHeader)
		packetsCount += 1
		if err == io.EOF {
			fmt.Println("Reached EOF")
			break
		} else if err != nil {
			panic(err)
		}

		if pcapPacketHeader.Length != pcapPacketHeader.UntruncLength {
			fmt.Print(packetsCount)
			panic("Length is truncated")
		}

		var pType ProtocolType
		err = binary.Read(r, binary.LittleEndian, &pType)
		if err != nil {
			panic(err)
		}

		if pType.ProtocolType != 2 {
			panic("protocol type wrong")
		}

		var ipHeader IpHeader
		err = binary.Read(r, binary.BigEndian, &ipHeader)
		if err != nil {
			panic(err)
		}

		if ipHeader.Length+4 != uint16(pcapPacketHeader.Length) {
			panic("ip header length wrong")
		}

		tcp := make([]byte, ipHeader.Length-20)
		_, err = io.ReadFull(r, tcp)

		sPort := binary.BigEndian.Uint16(tcp[:2])
		dPort := binary.BigEndian.Uint16(tcp[2:4])

		if sPort != 80 && dPort != 80 {
			panic("one of ports is not 80")
		}

		flags := tcp[13:14]

		ackFlag := flags[0] & 16 > 0
		synFlag := flags[0] & 2 > 0

		if sPort == 80 && ackFlag  {
			ackCount += 1
		}

		if dPort == 80 && synFlag {
			synCount += 1
		}

	}

	fmt.Printf("%v packets parsed\n", packetsCount)
	fmt.Printf("ackCount: %v, synCount %v\n", ackCount, synCount)
	fmt.Printf("%v is the percentage of acked syns\n", ackCount*100/synCount)
}
