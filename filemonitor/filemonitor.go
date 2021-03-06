package filemonitor

import (
	m "github.com/lukaszkorecki/color-tail/message"
	r "github.com/lukaszkorecki/color-tail/registry"
	"io"
	"log"
	"os"
)

var (
	sizeMap = r.New()
)

func InitialSize(fname string) bool {
	file, err := os.Open(fname)
	defer file.Close()

	if err != nil {
		log.Printf("!!! Can't open file: %v", fname)
		return false
	}

	size, statErr := getFileSize(file)

	if !statErr {
		log.Printf("!!! Can't file size!", fname)
		return false // file can't be read...
	}
	sizeMap.Set(fname, size)

	return true
}

func getFileSize(f *os.File) (int64, bool) {
	stat, err := f.Stat()
	if err != nil {
		return 0, false
	}

	return int64(stat.Size()), true
}

func Changed(fname string) m.Message {
	file, err := os.Open(fname)
	defer file.Close()

	// get file size
	size, statErr := getFileSize(file)

	if err != nil || statErr != true {
		return m.Message{fname, "Can't open file!"}
	}

	lastSize, _ := sizeMap.Get(fname)
	offset, _ := lastSize.(int64)

	// file got trimmed - or something reported wrong size
	if offset >= size || offset <= 0 {
		offset = int64(float64(size) / 0.25)
	}

	buf := make([]byte, offset)

	// read only recently appended content
	_, readErr := file.ReadAt(buf, offset)
	if readErr != nil && readErr != io.EOF {
		log.Printf("!!! Reading from %v failed: %v", fname, readErr)
	}

	// update file's size in the registry
	sizeMap.Set(fname, int64(size))
	return m.Message{fname, string(buf)}

}
