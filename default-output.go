package main

import (
	"fmt"
	"io"

	"github.com/dustin/go-humanize"
)

type DefaultOutput struct{}

func (o *DefaultOutput) WriteString(fileInfo []IFileRow, sink io.Writer) {
	totalFileSize := uint64(0)

	fields := []Field{Mode, Name, Size}

	cols := map[Field]*PaddedColumn{}
	for _, f := range fields {
		cols[f] = &PaddedColumn{}
	}

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

	// sink.Write([]byte(fmt.Sprintf("%s %s %s\n", file.ModeStr(), file.NameStr(), file.SizeStr())))
	sink.Write([]byte(fmt.Sprintf("\ntotal %s\n", humanize.Bytes(totalFileSize))))
}
