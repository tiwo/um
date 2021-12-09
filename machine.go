package main

import "os"

type Platter uint32
type Character uint8
type Array []Platter

type State struct {
	arrays    map[Platter]Array
	registers [8]Platter
	finger    Platter
}

func NewState() State {
	um := State{
		make(map[Platter]Array),
		[8]Platter{0, 0, 0, 0, 0, 0, 0, 0},
		0,
	}
	return um
}

func readFile(f os.File) (Array, error) {

}

func (um State) ReadProgramFromFile(f os.File) error {

}
