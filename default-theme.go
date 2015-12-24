package main

import "github.com/fatih/color"

type DefaultTheme struct{}

type DefaultThemeRow struct {
	IFileRow
}

var (
	c_blueBold    = color.New(color.FgBlue, color.Bold).SprintFunc()
	c_blue        = color.New(color.FgBlue).SprintFunc()
	c_whiteMedium = color.New().SprintFunc()
	c_whiteDark   = color.New(color.Faint).SprintFunc()
)

func (t *DefaultTheme) Format(s IFileRow) IFileRow {
	return &DefaultThemeRow{s}
}

func (r *DefaultThemeRow) Field(f Field) IFileField {
	n := r.IFileRow.Field(f)

	switch f {
	case Name:
		if r.IsDir() {
			return &FileField{c_blueBold(n.String()), n.Width()}
		} else {
			return &FileField{c_whiteMedium(n.String()), n.Width()}
		}

	case Mode:
		if r.IsDir() {
			return &FileField{c_blue(n.String()), n.Width()}
		} else {
			return &FileField{c_whiteDark(n.String()), n.Width()}
		}

	case Size:
		if r.IsDir() {
			return &FileField{c_blue(n.String()), n.Width()}
		} else {
			return &FileField{c_whiteDark(n.String()), n.Width()}
		}

	default:
		panic("Unimplemented case for Field")
	}
}
