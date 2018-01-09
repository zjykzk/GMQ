package store

import (
	"os"
	"unsafe"
)

type mappedfile struct {
	file          *os.File
	sz            int64
	data          []byte
	wrotePosition int64
	path          string
}

func newMappedfile(path string, sz int) (*mappedfile, error) {
	if path == "" {
		return nil, ErrBadPath
	}

	if sz <= 0 {
		return nil, ErrBadSize
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	b, err := mmap(int(file.Fd()), sz)
	if err != nil {
		return nil, err
	}

	return &mappedfile{
		path: path,
		file: file,
		data: b,
	}, nil
}

func (m *mappedfile) Write(b []byte) (c int, err error) {
	c, err = m.file.WriteAt(b, m.wrotePosition)
	if err != nil {
		return
	}
	m.wrotePosition += int64(c)
	return
}

func (m *mappedfile) close() error {
	return unmap(uintptr(unsafe.Pointer(&m.data[0])), uintptr(m.sz))
}

func (m *mappedfile) flush() error {
	return flush(uintptr(unsafe.Pointer(&m.data[0])), uintptr(m.wrotePosition))
}
