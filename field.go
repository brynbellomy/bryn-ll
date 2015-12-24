package main

type (
	IFileField interface {
		String() string
		Width() int
	}

	FileField struct {
		S string
		W int
	}
)

func (f *FileField) String() string {
	return f.S
}

func (f *FileField) Width() int {
	return f.W
}
