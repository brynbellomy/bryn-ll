package main

import (
	"fmt"
	"io"

	"github.com/dustin/go-humanize"
	"go.uber.org/multierr"
)

type DefaultOutput struct{}

func (o *DefaultOutput) WriteString(fileInfo []IFileRow, sink io.Writer) {
	var (
		totalFileSize = uint64(0)

		fields = []Field{Size, Name, Link, Mode}
		cols   = map[Field]*PaddedColumn{
			Mode: &PaddedColumn{Align: AlignLeft},
			Size: &PaddedColumn{Align: AlignRight},
			Name: &PaddedColumn{Align: AlignLeft},
			Link: &PaddedColumn{Align: AlignLeft},
		}
	)

	var merr error
	for _, file := range fileInfo {
		if !file.IsDir() {
			totalFileSize += uint64(file.Size())
		}
		for _, f := range fields {
			field, err := file.Field(f)
			cols[f].AddRow(field)
			merr = multierr.Append(merr, err)
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
	if merr != nil {
		sink.Write([]byte(fmt.Sprintf("\nErrors encountered:\n%+v", merr)))
	}
}
