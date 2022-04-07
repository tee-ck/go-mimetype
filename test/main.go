package main

import (
	_ "embed"
	"fmt"
)

func main() {
	fmt.Printf("%b\n", 0x00)
	fmt.Printf("%b\n", 0x09)
	fmt.Printf("%b\n", 0x0A)
	fmt.Printf("%b\n", 0x0C)
	fmt.Printf("%b\n", 0x0D)
	fmt.Printf("%b\n", 0x1C)
	fmt.Printf("%b\n", 0x20)
}
