package testdata

import (
	_ "embed"
)

//go:embed files/img.png
var PNG Uint8Array

//go:embed files/img.jpg
var JPEG Uint8Array

//go:embed files/img.heic
var HEIC Uint8Array

//go:embed files/audio.ogg
var OGG Uint8Array

//go:embed files/audio.m4a
var M4A Uint8Array

//go:embed files/video.mp4
var MP4 Uint8Array

//go:embed files/video-hevc.mp4
var HEVC Uint8Array

//go:embed files/video.webm
var WEBM Uint8Array

//go:embed files/document.pdf
var PDF Uint8Array

//go:embed files/document.docx
var DOCX Uint8Array

//go:embed files/document.xlsx
var XLSX Uint8Array

//go:embed files/document.pptx
var PPTX Uint8Array

//go:embed files/archive.rar
var RAR Uint8Array

//go:embed files/archive.zip
var ZIP Uint8Array

//go:embed files/text.json
var JSON Uint8Array
