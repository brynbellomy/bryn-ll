package main

import "strings"

type (
	PaddedColumn struct {
		items    []IFileField
		maxWidth int
		Align
	}

	Align int
)

const (
	AlignLeft Align = iota
	AlignRight
)

func (c *PaddedColumn) AddRow(row IFileField) {
	c.items = append(c.items, row)
}

func (c *PaddedColumn) Calculate() {
	c.maxWidth = c.MaxWidth()
}

func (c *PaddedColumn) MaxWidth() int {
	var max int
	for _, row := range c.items {
		if row.Width() > max {
			max = row.Width()
		}
	}
	return max
}

func (c *PaddedColumn) GetString(idx int) string {
	strlen := c.items[idx].Width()
	return c.Align.PadString(c.items[idx].String(), c.maxWidth-strlen)
}

func (a Align) PadString(s string, n int) string {
	switch a {
	case AlignLeft:
		return s + strings.Repeat(" ", n)
	case AlignRight:
		return strings.Repeat(" ", n) + s
	default:
		panic("Unknown case for Align")
	}
}
