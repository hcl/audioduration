package audioduration

import (
	"fmt"
	"io"
	"os"
)

// File Type Constant
const (
	TypeFlac int = 0
	TypeMp4  int = 1
	TypeMp3  int = 2
	TypeOgg  int = 3
	TypeDsd  int = 4
)

// Duration Get duration of specific music file type.
func Duration(file *os.File, filetype int) (float64, error) {
	var d float64 = 0
	var err error = nil
	file.Seek(0, io.SeekStart)
	switch filetype {
	case TypeFlac:
		d, err = FLAC(file)
	case TypeMp4:
		d, err = Mp4(file)
	case TypeMp3:
		d, err = Mp3(file)
	case TypeOgg:
		d, err = Ogg(file)
	case TypeDsd:
		d, err = DSD(file)
	default:
		err = fmt.Errorf("unsupported type: %d", filetype)
	}
	file.Seek(0, io.SeekStart)
	return d, err
}
