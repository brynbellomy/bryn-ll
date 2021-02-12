package main

import (
	"fmt"

	"github.com/fatih/color"
)

type DefaultTheme struct{}

type DefaultThemeRow struct {
	IFileRow
}

var (
	c_blueBold    = color.New(color.FgBlue, color.Bold).SprintFunc()
	c_blue        = color.New(color.FgBlue).SprintFunc()
	c_red         = color.New(color.FgRed).SprintFunc()
	c_whiteMedium = color.New().SprintFunc()
	c_whiteDark   = color.New(color.Faint).SprintFunc()
)

func (t *DefaultTheme) Format(s IFileRow) IFileRow {
	return &DefaultThemeRow{s}
}

func (r *DefaultThemeRow) Field(f Field) (IFileField, error) {
	n, err := r.IFileRow.Field(f)

	switch f {
	case Name:
		if r.IsDir() {
			return &FileField{c_blueBold(n.String()), n.Width(), false}, err
		} else {
			return &FileField{c_whiteMedium(n.String()), n.Width(), false}, err
		}

	case Link:
		if n.Errored() {
			return &FileField{c_red(n.String()), n.Width(), false}, err
		}
		return &FileField{c_blue(n.String()), n.Width(), false}, err

	case Mode:
		if r.IsDir() {
			return &FileField{c_blue(n.String()), n.Width(), false}, err
		} else {
			return &FileField{c_whiteDark(n.String()), n.Width(), false}, err
		}

	case Size:
		if r.IsDir() {
			return &FileField{c_blue(n.String()), n.Width(), false}, err
		} else {
			return &FileField{c_whiteDark(n.String()), n.Width(), false}, err
		}

	default:
		panic(fmt.Sprintf("Unimplemented case for Field: %v", f))
	}
}
