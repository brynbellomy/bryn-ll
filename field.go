package main

type (
	IFileField interface {
		String() string
		Width() int
		Errored() bool
	}

	FileField struct {
		S       string
		W       int
		errored bool
	}
)

func (f *FileField) String() string {
	return f.S
}

func (f *FileField) Width() int {
	return f.W
}

func (f *FileField) Errored() bool {
	return f.errored
}
