package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type BmpImage struct {
	width           uint32
	height          uint32
	startingAddress uint32
	bitsPerPixel    uint16
	rowSize         uint64
	pixelData       [][][]byte
}

func ReadBmpData(filename string) BmpImage {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	w := binary.LittleEndian.Uint32(file[0x12 : 0x12+4])
	h := binary.LittleEndian.Uint32(file[0x16 : 0x16+4])
	bpp := binary.LittleEndian.Uint16(file[0x1c : 0x1c+2])
	startingAddress := binary.LittleEndian.Uint32(file[0xA : 0xA+4])

	bytesPerPixel := bpp >> 3

	rowSize := uint64(bpp) * uint64(w) >> 3
	if rowSize%4 != 0 {
		rowSize = rowSize + (4 - rowSize%4)
	}

	pixelData := make([][][]byte, h)

	for i := uint32(0); i < h; i++ {
		pixelData[i] = make([][]byte, w)
		for j := uint32(0); j < w; j++ {
			pixelPosition := int(startingAddress) + int(i)*int(rowSize) + int(j)*int(bytesPerPixel)
			pixelData[i][j] = file[pixelPosition : pixelPosition+int(bytesPerPixel)]
		}

	}

	return BmpImage{width: w, height: h, bitsPerPixel: bpp, startingAddress: startingAddress, rowSize: rowSize, pixelData: pixelData}
}

func main() {
  fmt.Print(os.Args[1])
	bmpImage := ReadBmpData(os.Args[1])
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rotatedHeight := bmpImage.width
	rotatedWidth := bmpImage.height
	rotatedRowSize := uint64(bmpImage.bitsPerPixel) * uint64(rotatedWidth) >> 3
	var padding uint64
	if rotatedRowSize%4 != 0 {
		padding = (4 - rotatedRowSize%4)
		rotatedRowSize = rotatedRowSize + padding
	}

	rotatedPixelData := make([][][]byte, rotatedHeight)

	for i := uint32(0); i < rotatedHeight; i++ {
		rotatedPixelData[i] = make([][]byte, rotatedWidth)
	}

	for i := uint32(0); i < rotatedHeight; i++ {
		for j := uint32(0); j < rotatedWidth; j++ {
			rotatedPixelData[i][j] = bmpImage.pixelData[j][rotatedHeight-1-i]
		}
	}

	newImage := make([]byte, bmpImage.startingAddress)
	copy(newImage, file[:bmpImage.startingAddress])

	for i := (0); i < len(rotatedPixelData); i++ {
		for j := (0); j < len(rotatedPixelData[0]); j++ {
			for _, v := range rotatedPixelData[i][j] {
				newImage = append(newImage, v)
			}
		}

		for i := uint8(0); i < uint8(padding); i++ {
			p := byte(0)
			newImage = append(newImage, p)
		}
	}

	w := file[0x12 : 0x12+4]

	var totalSize []byte
	totalSize = binary.LittleEndian.AppendUint64(totalSize, uint64(len(newImage[bmpImage.startingAddress:])))

	for i := range w {
		newImage[0x12+i] = file[0x16+i]
		newImage[0x16+i] = file[0x12+i]
		newImage[0x22+i] = totalSize[i]
	}

	err = os.WriteFile("./teste.bmp", newImage, 0644)
	if err != nil {
		panic(err)
	}
}
