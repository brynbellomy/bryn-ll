package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
)

type (
	DefaultRowFormatter struct{}

	DefaultFileInfo struct {
		dir string
		os.FileInfo
	}
)

func (f *DefaultRowFormatter) Format(dir string, row os.FileInfo) IFileRow {
	return &DefaultFileInfo{dir, row}
}

func (f *DefaultFileInfo) Field(field Field) (IFileField, error) {
	switch field {
	case Name:
		n := f.FileInfo.Name()
		return &FileField{n, len(n), false}, nil

	case Link:
		if f.FileInfo.Mode()&os.ModeSymlink != 0 {
			fullpath := filepath.Join(f.dir, f.FileInfo.Name())
			link, err := os.Readlink(fullpath)
			if err != nil {
				return &FileField{"", 0, true}, fmt.Errorf("error following symlink: %v", err)
			}
			if !filepath.IsAbs(link) {
				link = filepath.Join(f.dir, link)
				link, err = filepath.Abs(link)
				if err != nil {
					return &FileField{link, 0, true}, err
				}
			}
			relpath, err := filepath.Rel(f.dir, link)
			if err != nil {
				return &FileField{link, 0, true}, fmt.Errorf("error getting symlink relative path: %v", err)
			}
			s := "-> " + relpath

			_, err = os.Stat(link)
			if os.IsNotExist(err) {
				return &FileField{s, len(s), true}, nil
			}
			return &FileField{s, len(s), false}, nil
		}
		return &FileField{"", 0, false}, nil

	case Mode:
		permBits := f.FileInfo.Mode().Perm()
		n := fmt.Sprintf("%o", permBits)
		return &FileField{n, len(n), false}, nil

	case Size:
		s := f.FileInfo.Size()
		n := humanize.Bytes(uint64(s))
		return &FileField{n, len(n), false}, nil

	default:
		panic(fmt.Sprintf("Unimplemented case for Field: %v", field))
	}
}
