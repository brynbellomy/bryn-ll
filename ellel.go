package main

import (
	"io"
	"io/ioutil"
	"os"
	"sort"
)

type (
	Ellel struct {
		Dir       string
		Formatter IRowFormatter
		Theme     ITheme
		Output    IOutput
	}

	IRowFormatter interface {
		Format(dir string, row os.FileInfo) IFileRow
	}

	ITheme interface {
		Format(s IFileRow) IFileRow
	}

	IFileRow interface {
		os.FileInfo
		Field(f Field) (IFileField, error)
	}

	IOutput interface {
		WriteString(fileInfo []IFileRow, sink io.Writer)
	}

	Field int
)

const (
	Name Field = iota
	Link
	Mode
	Size
)

func (el *Ellel) Render() error {
	files, err := ioutil.ReadDir(el.Dir)
	if err != nil {
		return err
	}

	sort.Sort(Files(files))

	fs := make([]IFileRow, len(files))
	for i, file := range files {
		fs[i] = el.Formatter.Format(el.Dir, file)
		fs[i] = el.Theme.Format(fs[i])
	}

	el.Output.WriteString(fs, os.Stdout)
	return nil
}
