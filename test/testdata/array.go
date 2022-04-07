package testdata

import (
	"encoding/hex"
	"fmt"
)

type Uint8Array []byte

func (u Uint8Array) Hex() string {
	return hex.EncodeToString(u)
}

func (u Uint8Array) HexArray() HexArray {
	return HexArray(u)
}

func (u Uint8Array) String() string {
	return string(u)
}

type HexArray []byte

func (h HexArray) Indexed() IndexedHexArray {
	return IndexedHexArray(h)
}

func (h HexArray) PrintTable(row, col int) {
	fmt.Printf("      %4d", 0)
	for i := 1; i < col; i++ {
		fmt.Printf(" %4d", i)
	}
	fmt.Println()

	for i := 0; i < row; i++ {
		fmt.Printf("%4d ", i*col)
		fmt.Printf("[")

		fmt.Printf("0x%02x", h[(i*col)])
		for j := 1; j < col; j++ {
			fmt.Printf(" 0x%02x", h[(i*col)+j])
		}
		fmt.Printf("]\n")
	}
}

func (h HexArray) String() string {
	s := make([]string, 0, len(h))
	for _, v := range h {
		s = append(s, fmt.Sprintf("0x%02x", v))
	}

	return fmt.Sprint(s)
}

type IndexedHexArray []byte

func (i IndexedHexArray) String() string {
	s := make([]string, 0, len(i))
	for i, v := range i {
		s = append(s, fmt.Sprintf("%d:0x%02x", i, v))
	}

	return fmt.Sprint(s)
}
