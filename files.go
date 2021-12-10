package main

import (
	"errors"
	"fmt"
	"io"
)

func readArrayFromFile(f io.Reader) (Array, error) {
	var countToFour = 0
	var currentPlatter Platter = 0
	result := make([]Platter, 0)
	buf := make([]byte, 3)
	var bytesRead int
	var err error
	for err = nil; err == nil; bytesRead, err = f.Read(buf) {
		for i := 0; i < bytesRead; i++ {
			currentPlatter = (currentPlatter << 8) + Platter(buf[i])
			countToFour++
			if countToFour == 4 {
				result = append(result, currentPlatter)
				currentPlatter = 0
				countToFour = 0
			}
		}
	}

	if err != nil {
		if !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrUnexpectedEOF) {
			return nil, err
		}
		// EOF and ErrUnexpectedEOF are okay - ignore them here
	}

	if countToFour != 0 {
		fmt.Println("io.ErrUnexpectedEOF")
		return nil, io.ErrUnexpectedEOF
	}

	return result, nil

}

func ReadProgramFromFile(f io.Reader) (State, error) {
	um := NewState()

	um.arrays[0], err := readArrayFromFile(f)
	if err != nil {
		return State{}, err
	}
	return um, nil
}
