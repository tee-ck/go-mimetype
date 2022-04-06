package main

import (
	_ "embed"
	"fmt"
)

//go:embed testdata/img.png
var PNG []byte

//go:embed testdata/img.jpg
var JPEG []byte

//go:embed testdata/video.mp4
var MP4 []byte

//go:embed testdata/document.pdf
var PDF []byte

//go:embed testdata/img.heic
var HEIC []byte

func toHexSlice(buf []byte) []string {
	items := make([]string, 0, len(buf))

	for i, b := range buf {
		items = append(items, fmt.Sprintf("%d:0x%02X", i, b))
	}

	return items
}

func main() {
	//fmt.Println(hex.EncodeToString([]byte("<?xml")))
	//fmt.Println(hex.EncodeToString([]byte("<!DOCTYPE HTML")))
	//fmt.Println(hex.EncodeToString([]byte("{")))
	//fmt.Println(hex.EncodeToString([]byte("d8:announce")))
	//fmt.Println(to_hex_slice(MP4[:32]))
	//fmt.Println(to_hex_slice(HEIC[:32]))
	fmt.Println(toHexSlice(HEIC[3:32]))
	fmt.Println(toHexSlice([]byte("heic")))
	fmt.Println(toHexSlice([]byte("ftyp")))
	fmt.Println(string(HEIC[3:32]))
}
