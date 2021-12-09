package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.um")
	if err != nil {
		log.Fatal(err)
	}

	x, err := readFile(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("... %#v\n", x)
}
