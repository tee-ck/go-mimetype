package main

import (
	_ "embed"
	"fmt"
)

//go:embed testdata/files/audio.m4a
var audio []byte

func main() {
	for i, e := range audio[:48] {
		fmt.Printf("%02d: %02X [%c]\n", i, e, rune(e))
	}
}
