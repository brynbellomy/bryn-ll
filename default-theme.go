package main

import "github.com/fatih/color"

type DefaultTheme struct{}

type DefaultThemeRow struct {
	IFileRow
}

var (
	c_dirBold = color.New(color.FgBlue, color.Bold).SprintFunc()
	c_dir     = color.New(color.FgBlue).SprintFunc()

	c_dark = color.New(color.Faint).SprintFunc()
)

func (t *DefaultTheme) Format(s IFileRow) IFileRow {
	return &DefaultThemeRow{s}
}

func (r *DefaultThemeRow) Field(f Field) IFileField {
	n := r.IFileRow.Field(f)

	switch f {
	case Name:
		if r.IsDir() {
			return &FileField{c_dirBold(n.String()), n.Width()}
		} else {
			return r.IFileRow.Field(f)
		}

	case Mode:
		if r.IsDir() {
			return &FileField{c_dir(n.String()), n.Width()}
		} else {
			return &FileField{c_dark(n.String()), n.Width()}
		}

	case Size:
		if r.IsDir() {
			return &FileField{c_dir(n.String()), n.Width()}
		} else {
			return &FileField{c_dark(n.String()), n.Width()}
		}

	default:
		panic("Unimplemented case for Field")
	}
}
