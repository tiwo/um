package main

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
