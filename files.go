package main

import (
	"io"
	"os"
)

func readFile(f os.File) (Array, error) {
	var countToFour = 0
	var currentPlatter Platter = 0
	result := make([]Platter)
	var buf [1024]byte
	var bytesRead int
	var err error
	for err = nil; err == nil; bytesRead, err := io.ReadFull(f, buf) {
		for i := 0; i < bytesRead; i++ {
			currentPlatter = (currentPlatter<<8) + uint8(buf[i])
			countToFour++
			if countToFour = 3 {
				append(result, currentPlatter)
				currentPlatter = 0
				countToFour = 0
			}
		}
	}

	switch err = err.(type) {
	case io.EOF:
		// noop
	case io.ErrUnexpectedEOF:
		// noop
	default:
		return nil, err
	}

	if countToFour != 0 {
		return nil, io.ErrUnexpectedEOF("the last platter was not received (file size should be divisible by 4 bytes")
	}

	return result, nil

}

func (um State) ReadProgramFromFile(f os.File) error {
	array, err := readFile(f)
	if err != nil {
		
	}

}
