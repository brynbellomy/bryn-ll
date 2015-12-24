package main

import (
	"fmt"
	"io"

	"github.com/dustin/go-humanize"
)

type DefaultOutput struct{}

func (o *DefaultOutput) WriteString(fileInfo []IFileRow, sink io.Writer) {
	var (
		totalFileSize = uint64(0)

		fields = []Field{Size, Name, Mode}
		cols   = map[Field]*PaddedColumn{
			Mode: &PaddedColumn{Align: AlignLeft},
			Size: &PaddedColumn{Align: AlignRight},
			Name: &PaddedColumn{Align: AlignLeft},
		}
	)

	for _, file := range fileInfo {
		if !file.IsDir() {
			totalFileSize += uint64(file.Size())
		}
		for _, f := range fields {
			cols[f].AddRow(file.Field(f))
		}
	}

	for _, f := range fields {
		cols[f].Calculate()
	}

	sink.Write([]byte("\n"))

	for i := range fileInfo {
		for _, f := range fields {
			str := cols[f].GetString(i)
			sink.Write([]byte(str + " "))
		}
		sink.Write([]byte("\n"))
	}

	sink.Write([]byte(fmt.Sprintf("\ntotal %s\n", humanize.Bytes(totalFileSize))))
}
