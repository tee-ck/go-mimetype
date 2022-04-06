package testdata

import (
	_ "embed"
)

//go:embed img.png
var PNG []byte

//go:embed img.jpg
var JPEG []byte

//go:embed video.mp4
var MP4 []byte

//go:embed document.pdf
var PDF []byte

//go:embed img.heic
var HEIC []byte
