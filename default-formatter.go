package main

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
)

type (
	DefaultRowFormatter struct{}

	DefaultFileInfo struct {
		os.FileInfo
	}
)

func (f *DefaultRowFormatter) Format(row os.FileInfo) IFileRow {
	return &DefaultFileInfo{row}
}

func (f *DefaultFileInfo) Field(field Field) IFileField {
	switch field {
	case Name:
		n := f.FileInfo.Name()
		return &FileField{n, len(n)}

	case Mode:
		permBits := f.FileInfo.Mode().Perm()
		n := fmt.Sprintf("%o", permBits)
		return &FileField{n, len(n)}

	case Size:
		s := f.FileInfo.Size()
		n := humanize.Bytes(uint64(s))
		return &FileField{n, len(n)}

	default:
		panic("Unimplemented case for Field")
	}
}
