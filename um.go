package main

type Platter uint32
type Array []Platter

type State struct {
	arrays    map[Platter]Array
	registers [8]Platter
}

func main() {

}
