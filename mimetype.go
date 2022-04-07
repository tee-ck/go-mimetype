package mimetype

import (
	"bytes"
	_ "embed"
)

func Detect(b []byte) (mimetype string) {
	if b[0] < 0x21 {
		for x := 0; x < len(b); x++ {
			if b[x] == 0x00 || b[x] == 0x09 || b[x] == 0x0A || b[x] == 0x0C || b[x] == 0x0D || b[x] == 0x1C || b[x] == 0x20 {
				continue
			}

			if x > 0 {
				b = b[x:]
			}
			break
		}
	}

	switch b[0] {
	case 0x01:
		if b[1] == 0x00 {
			if b[2] == 0x00 && b[3] == 0x00 {
				return "font/tff"
			}

			return "image/x-icon"
		}
	case 0x02:
		if b[1] == 0x00 {
			return "image/x-icon"
		}
	case 0x18:
		if b[1] == 0x66 && b[2] == 0x74 && b[3] == 0x79 && b[4] == 0x70 {
			switch b[5] {
			case 0x68:
				if b[6] == 0x65 && b[7] == 0x69 && b[8] == 0x63 {
					return "image/heic"
				}
			default:
				if b[17] == 0x68 && b[18] == 0x65 && b[19] == 0x69 && b[20] == 0x63 {
					return "image/heic"
				}
			}

			return "video/mp4"
		}
	case 0x1A:
		if b[1] == 0x45 && b[2] == 0xDF && b[3] == 0xA3 {
			if b[4] == 0x9F {
				return "video/webm"
			}

			// HEVC video
			return "video/mp4"
		}
	case 0x1F:
		if b[1] == 0x8B {
			return "application/gzip"
		}
	case 0x21:
		if b[1] == 0x3c && b[2] == 0x61 && b[3] == 0x72 && b[4] == 0x63 && b[5] == 0x68 && b[6] == 0x3e && b[7] == 0x0A {
			return "application/x-archive"
		}
	case 0x23:
		if b[1] == 0x21 && b[2] == 0x41 && b[3] == 0x4D && b[4] == 0x52 {
			return "audio/amr"
		}
	case 0x25:
		switch b[1] {
		case 0x21:
			if b[2] == 0x50 && b[3] == 0x53 && b[4] == 0x2D && b[5] == 0x41 && b[6] == 0x64 && b[7] == 0x6F {
				return "application/postscript"
			}
		case 0x50:
			if b[2] == 0x44 && b[3] == 0x46 && b[4] == 0x2D {
				return "application/pdf"
			}
		}
	case 0x2E:
		switch b[1] {
		case 0x52:
			if b[2] == 0x4D && b[3] == 0x46 {
				return "application/vnd.rn-realmedia-vbr"
			}
		case 0x73:
			if b[2] == 0x6E && b[3] == 0x64 {
				return "audio/basic"
			}
		}
	case 0x30:
		if b[1] == 0x26 && b[2] == 0xB2 && b[3] == 0x75 && b[4] == 0x8E && b[5] == 0x66 && b[6] == 0xCF && b[7] == 0x11 {
			return "video/x-ms-wmv"
		}
	case 0x3C:
		switch b[1] {
		case 0x21:
			b = bytes.ToUpper(b)
			if bytes.Compare(b[2:16], []byte("<!DOCTYPE HTML")) == 0 {
				return "text/html"
			}
		case 0x3F:
			b = bytes.ToUpper(b)
			if bytes.Index(b, []byte("<SVG")) >= 0 || bytes.Index(b, []byte("<!DOCTYPE SVG")) >= 0 {
				return "image/svg+xml"

			} else if b[2] == 0x78 && b[3] == 0x6D && b[4] == 0x6C {
				return "text/xml"

			} else if bytes.Index(b, []byte("<?PHP")) > -1 {
				return "application/x-php"
			}
		}
	case 0x37:
		switch b[1] {
		case 0x7A:
			if b[2] == 0xBC && b[3] == 0xAF && b[4] == 0x27 && b[5] == 0x1C {
				return "application/x-7z-compressed"
			}
		case 0xA4:
			if b[3] == 0x30 && b[4] == 0xEC {
				return "application/zstd"
			}
		}
	case 0x38:
		if b[1] == 0x42 && b[2] == 0x50 && b[3] == 0x53 {
			return "image/vnd.adobe.photoshop"
		}
	case 0x42:
		switch b[1] {
		case 0x4C:
			if b[2] == 0x45 && b[3] == 0x4E && b[4] == 0x44 && b[5] == 0x45 && b[6] == 0x52 {
				return "application/x-blender"
			}

		case 0x4D:
			return "image/bmp"

		case 0x50:
			if b[2] == 0x47 && b[3] == 0x46 {
				return "image/bpg"
			}
		case 0x5A:
			if b[2] == 0x68 {
				return "application/x-bzip2"
			}
		}
	case 0x43:
		switch b[1] {
		case 0x57:
			if b[2] == 0x53 {
				return "application/x-cfs-compressed"
			}
		case 0x72:
			if b[2] == 0x32 && b[3] == 0x34 {
				return "application/x-chrome-extension"
			}
		}
	case 0x46:
		switch b[1] {
		case 0x4C:
			switch b[2] {
			case 0x49:
				if b[3] == 0x46 {
					return "image/flif"
				}
			case 0x4C:
				if b[3] == 0x56 {
					return "video/flv"
				}
			}
		case 0x4F:
			if b[2] == 0x52 && b[3] == 0x4D && b[8] == 0x41 && b[9] == 0x49 && b[10] == 0x46 && b[11] == 0x46 {
				return "audio/aiff"
			}
		case 0x57:
			if b[2] == 0x53 {
				return "application/x-shockwave-flash"
			}
		}
	case 0x47:
		if b[1] == 0x49 && b[2] == 0x46 && b[3] == 0x38 {
			switch b[4] {
			case 0x37, 0x39:
				if b[5] == 0x61 {
					return "image/gif"
				}
			}
		}
	case 0x49:
		if b[1] == 0x49 {
			switch b[2] {
			case 0x2A:
				if b[3] == 0x00 {
					if b[4] == 0x10 && b[5] == 0x00 && b[6] == 0x00 && b[7] == 0x00 && b[8] == 0x43 && b[9] == 0x52 {
						return "image/x-canon-cr2"
					}

					return "image/tiff"
				}
			case 0x44:
				if b[3] == 0x43 {
					return "audio/mp3"
				}
			}
		}
	case 0x4C:
		if b[1] == 0x5A && b[2] == 0x49 && b[3] == 0x50 {
			return "application/x-lzip"
		}
	case 0x4D:
		switch b[1] {
		case 0x4C:
			if b[2] == 0x56 && b[3] == 0x49 {
				return "video/mlv"
			}
		case 0x4D:
			if b[2] == 0x00 && (b[3] == 0x2A || b[4] == 0x2B) {
				return "image/tiff"
			}
		case 0x53:
			if b[2] == 0x43 && b[3] == 0x46 {
				return "application/x-ms-wim"
			}
		case 0x54:
			if b[2] == 0x68 && b[3] == 0x64 {
				return "audio/midi"
			}
		case 0x5A:
			return "application/x-msdos-program"
		}
	case 0x4F:
		switch b[1] {
		case 0x54:
			if b[2] == 0x54 && b[3] == 0x4F {
				return "font/otf"
			}
		case 0x67:
			if b[2] == 0x67 && b[3] == 0x53 {
				return "audio/ogg"
			}
		}
	case 0x50:
		if b[1] == 0x4B {
			switch b[2] {
			case 0x03:
				if b[3] == 0x04 {
					switch b[4] {
					case 0x14:
						if b[5] == 0x00 && b[6] == 0x08 && b[7] == 0x00 && b[8] == 0x08 && b[9] == 0x00 {
							return "application/java-archive"
						}
					case 0x0A:
						if b[5] == 0x00 && b[6] == 0x02 && b[7] == 0x00 {
							return "application/epub"
						}
					}

					return "application/x-zip"
				}
			case 0x05:
				if b[3] == 0x05 {
					return "application/x-zip-compressed"
				}
			case 0x07:
				if b[3] == 0x08 {
					return "application/x-7z-compressed"
				}
			}
		}
	case 0x52:
		switch b[1] {
		case 0x49:
			if b[2] == 0x46 && b[3] == 0x46 {
				switch b[8] {
				case 0x57:
					switch b[9] {
					case 0x41:
						if b[10] == 0x56 && b[11] == 0x45 {
							return "audio/wav"
						}
					case 0x45:
						if b[10] == 0x42 && b[11] == 0x50 {
							return "image/webp"
						}
					}
				case 0x41:
					if b[9] == 0x56 && b[10] == 0x49 && b[11] == 0x20 {
						return "video/avi"
					}
				}
			}
		case 0x61:
			if b[2] == 0x72 && b[3] == 0x21 && b[4] == 0x1A && b[5] == 0x07 {
				if b[6] == 0x00 || (b[6] == 0x01 && b[7] == 0x00) {
					return "application/x-rar-compressed"
				}
			}
		}
	case 0x5A:
		if b[1] == 0x4D {
			return "application/x-msdos-program"
		}
	case 0x5F:
		if b[1] == 0x27 && b[2] == 0xA8 && b[3] == 0x89 {
			return "application/java-archive"
		}
	case 0x61:
		if b[1] == 0x73 && b[2] == 0x6D {
			return "application/wasm"
		}
	case 0x6A:
		if b[1] == 0x50 && b[2] == 0x20 {
			return "image/jpeg"
		}
	case 0x64:
		if b[1] == 0x38 && b[2] == 0x3A && b[3] == 0x61 && b[4] == 0x6E && b[5] == 0x6E && b[6] == 0x6F && b[7] == 0x75 && b[8] == 0x6E && b[9] == 0x63 && b[10] == 0x65 {
			return "application/x-bittorrent"
		}
	case 0x66:
		switch b[1] {
		case 0x4C:
			if b[2] == 0x61 && b[3] == 0x43 {
				return "audio/flac"
			}
		case 0x74:
			if b[2] == 0x79 && b[3] == 0x70 {
				switch b[4] {
				case 0x33:
					if b[5] == 0x67 && b[6] == 0x70 {
						return "video/3gpp"
					}
				case 0x4D:
					if b[5] == 0x34 && b[6] == 0x41 {
						return "audio/mp4"
					}
				}

				return "video/mp4"
			}
		}
	case 0x6B:
		if b[1] == 0x6F && b[2] == 0x6C && b[3] == 0x79 {
			return "application/x-apple-diskimage"
		}
	case 0x6D:
		if b[1] == 0x6F && b[2] == 0x6F && b[3] == 0x76 {
			return "video/mov"
		}
	case 0x75:
		if b[1] == 0x73 && b[2] == 0x74 && b[3] == 0x61 && b[4] == 0x72 {
			if (b[5] == 0x00 && b[6] == 0x30 && b[7] == 0x30) || (b[5] == 0x20 && b[6] == 0x20 && b[7] == 0x00) {
				return "application/x-tar"
			}
		}
	case 0x77:
		if b[1] == 0x4F && b[2] == 0x46 {
			if b[3] == 0x46 {
				return "font/woff"

			} else if b[3] == 0x32 {
				return "font/woff2"
			}
		}
	case 0x78:
		if b[1] == 0x61 && b[2] == 0x72 && b[3] == 0x21 {
			return "application/x-xar"
		}
	case 0x7B:
		if b[1] == 0x5C && b[2] == 0x72 && b[3] == 0x74 && b[4] == 0x66 {
			return "text/rtf"
		}

		return "application/json"
	case 0x89:
		switch b[1] {
		case 0x50:
			if b[2] == 0x4e && b[3] == 0x47 {
				return "image/png"
			}
		case 0x89:
			if b[2] == 0x4E && b[3] == 0x47 && b[4] == 0x0D && b[5] == 0x0A && b[6] == 0x1A && b[7] == 0x0A {
				return "image/png"
			}
		}
	case 0xCA:
		if b[1] == 0xFE && b[2] == 0xBA && b[3] == 0xBE {
			return "application/java-vm"
		}
	case 0xD0:
		if b[1] == 0xCF && b[2] == 0x11 && b[3] == 0xE0 && b[4] == 0xA1 && b[5] == 0xB1 && b[6] == 0x1A && b[7] == 0xE1 {
			if len(b) > 524 {
				switch b[512] {
				case 0xA0:
					if b[513] == 0x46 {
						return "application/vnd.ms-powerpoint"
					}
				case 0xEC:
					if b[513] == 0xA5 {
						return "application/msword"
					}
				case 0x09:
					if b[513] == 0x08 {
						return "application/vnd.ms-excel"
					}
				}
			}

			return "application/msword"
		}
	case 0xDF:
		if b[1] == 0xBF && b[2] == 0x34 && b[3] == 0xEB && b[4] == 0xCE {
			return "application/pdf"
		}
	case 0xED:
		if b[1] == 0xAB && b[2] == 0xEE && b[3] == 0xDB {
			return "application/x-rpm"
		}
	case 0xFD:
		if b[1] == 0x37 && b[2] == 0x7A && b[3] == 0x58 && b[4] == 0x5A && b[5] == 0x00 {
			return "application/x-xz"
		}
	case 0xFF:
		switch b[1] {
		case 0xF1, 0xF9:
			return "audio/aac"
		case 0xFB, 0xF3, 0xF2:
			return "audio/mp3"
		case 0x4F:
			if b[2] == 0xFF && b[3] == 0x51 {
				return "image/jpeg"
			}
		case 0xD8:
			if b[2] == 0xFF && (b[3] == 0xD8 || b[3] == 0xE0 || b[3] == 0xE1 || b[3] == 0xE8 || b[3] == 0xEE) {
				return "image/jpeg"
			}
		}
	}

	return "application/octet-stream"
}
