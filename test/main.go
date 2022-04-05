package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(hex.EncodeToString([]byte("<?xml")))
	fmt.Println(hex.EncodeToString([]byte("<!DOCTYPE HTML")))
	fmt.Println(hex.EncodeToString([]byte(".snd")))
}
