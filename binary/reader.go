// Package binary provides abstraction for working with binary files.
package binary

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

// Reader interface for binary files
type Reader interface {
	Double() (float64, error)
	Discard(int) (int, error)
	Float() (float32, error)
	Int8() (int8, error)
	Int16() (int16, error)
	Int32() (int32, error)
	Int64() (int64, error)
	String() (string, error)
	UInt8() (uint8, error)
	UInt16() (uint16, error)
	UInt32() (uint32, error)
	UInt64() (uint64, error)
}

// File is os.File extension implementing Reader interface
type File struct {
	os.File
}

// OpenFile opens file at given path for reading. If there is an error, it
// will be of type *PathError
func OpenFile(path string) (*File, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &File{*fh}, nil
}

// Double reads float64 from File
func (fh *File) Double() (float64, error) {
	var out float64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Discard n bytes
func (fh *File) Discard(n int) (int, error) {
	m, err := fh.Seek(int64(n), 1)
	return int(m), err
}

// Float reads float32 from File
func (fh *File) Float() (float32, error) {
	var out float32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int8 reads int8 from File
func (fh *File) Int8() (int8, error) {
	var out int8
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int16 reads int16 from File
func (fh *File) Int16() (int16, error) {
	var out int16
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int32 reads int32 from File
func (fh *File) Int32() (int32, error) {
	var out int32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int64 reads int64 from File
func (fh *File) Int64() (int64, error) {
	var out int64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// String reads string from File
// Function first reads uint16 which denotes string length N, then N bytes are
// read and returned as string
func (fh *File) String() (string, error) {
	var strLen uint16
	if err := binary.Read(fh, binary.LittleEndian, &strLen); err != nil {
		return "", err
	}
	out := make([]byte, strLen, strLen)
	if _, err := fh.Read(out); err != nil {
		return "", err
	}
	return string(out), nil
}

// UInt8 reads uint8 from File
func (fh *File) UInt8() (uint8, error) {
	var out uint8
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt16 reads uint16 from File
func (fh *File) UInt16() (uint16, error) {
	var out uint16
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt32 reads uint32 from File
func (fh *File) UInt32() (uint32, error) {
	var out uint32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt64 reads uint64 from File
func (fh *File) UInt64() (uint64, error) {
	var out uint64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// BufferedFile extends bufio.Reader implementing Reader interface
type BufferedFile struct {
	bufio.Reader
}

// OpenBufferedFile opens file at given path for reading. If there is an
// error, it will be of type *PathError
func OpenBufferedFile(path string) (*BufferedFile, error) {
	fh, err := OpenFile(path)
	if err != nil {
		return nil, err
	}
	bufFh := new(BufferedFile)
	bufFh.Reset(fh)
	return bufFh, nil
}

// Double reads float64 from BufferedFile
func (fh *BufferedFile) Double() (float64, error) {
	var out float64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Discard n bytes
func (fh *BufferedFile) Discard(n int) (int, error) {
	// TODO: fixme, Discarding / bath reading was bugged
	return 0, fmt.Errorf("Not implemented")
}

// Float reads float32 from BufferedFile
func (fh *BufferedFile) Float() (float32, error) {
	var out float32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int8 reads int8 from BufferedFile
func (fh *BufferedFile) Int8() (int8, error) {
	var out int8
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int16 reads int16 from BufferedFile
func (fh *BufferedFile) Int16() (int16, error) {
	var out int16
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int32 reads int32 from BufferedFile
func (fh *BufferedFile) Int32() (int32, error) {
	var out int32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int64 reads int64 from BufferedFile
func (fh *BufferedFile) Int64() (int64, error) {
	var out int64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

func (fh *BufferedFile) String() (string, error) {
	var strLen uint16
	if err := binary.Read(fh, binary.LittleEndian, &strLen); err != nil {
		return "", err
	}
	out := make([]byte, strLen, strLen)
	if _, err := fh.Read(out); err != nil {
		return "", err
	}
	return string(out), nil
}

// UInt8 reads uint8 from BufferedFile
func (fh *BufferedFile) UInt8() (uint8, error) {
	var out uint8
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt16 reads uint16 from BufferedFile
func (fh *BufferedFile) UInt16() (uint16, error) {
	var out uint16
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt32 reads uint32 from BufferedFile
func (fh *BufferedFile) UInt32() (uint32, error) {
	var out uint32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt64 reads uint64 from BufferedFile
func (fh *BufferedFile) UInt64() (uint64, error) {
	var out uint64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}
