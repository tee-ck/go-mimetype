package mimetype

import (
	_ "embed"
	"testing"
)

var (
	//go:embed test/testdata/files/audio.m4a
	audioM4A []byte
)

func TestAudioM4A(t *testing.T) {
	t.Log(Detect(audioM4A))
}

//func BenchmarkJPEG(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(imageJPEG)
//	}
//}
//
//func BenchmarkNetJPEG(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(imageJPEG)
//	}
//}
//
//func BenchmarkPNG(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.PNG)
//	}
//}
//
//func BenchmarkNetPNG(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.PNG)
//	}
//}
//
//func BenchmarkHEIC(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.HEIC)
//	}
//}
//
//func BenchmarkNetHEIC(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.HEIC)
//	}
//}
//
//func BenchmarkOGG(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.OGG)
//	}
//}
//
//func BenchmarkNetOGG(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.OGG)
//	}
//}
//
//func BenchmarkM4A(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.M4A)
//	}
//}
//
//func BenchmarkNetM4A(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.M4A)
//	}
//}
//
//func BenchmarkMP4(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.MP4)
//	}
//}
//
//func BenchmarkNetMP4(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.MP4)
//	}
//}
//
//func BenchmarkHEVC(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.HEVC)
//	}
//}
//
//func BenchmarkNetHEVC(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.HEVC)
//	}
//}
//
//func BenchmarkWEBM(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.WEBM)
//	}
//}
//
//func BenchmarkNetWEBM(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.WEBM)
//	}
//}
//
//func BenchmarkPDF(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.PDF)
//	}
//}
//
//func BenchmarkNetPDF(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.PDF)
//	}
//}
//
//func BenchmarkRAR(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.RAR)
//	}
//}
//
//func BenchmarkNetRAR(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.RAR)
//	}
//}
//
//func BenchmarkZIP(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.ZIP)
//	}
//}
//
//func BenchmarkNetZIP(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.ZIP)
//	}
//}
//
//func BenchmarkJSON(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Detect(testdata.JSON)
//	}
//}
//
//func BenchmarkNetJSON(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		http.DetectContentType(testdata.JSON)
//	}
//}
