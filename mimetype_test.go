package mimetype

import (
	_ "embed"
	"net/http"
	"testing"
)

func BenchmarkJPG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect(JPEG)
	}
}

func BenchmarkNetJPG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.DetectContentType(JPEG)
	}
}

func BenchmarkPNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect(PNG)
	}
}

func BenchmarkNetPNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.DetectContentType(PNG)
	}
}

func BenchmarkMP4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect(MP4)
	}
}

func BenchmarkNetMP4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.DetectContentType(MP4)
	}
}

func BenchmarkPDF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect(PDF)
	}
}

func BenchmarkNetPDF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.DetectContentType(PDF)
	}
}
