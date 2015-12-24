package main

import (
	"os"
	"sort"
)

type Files []os.FileInfo

func (f Files) Len() int {
	return len(f)
}

func (f Files) Less(i, j int) bool {
	lhs := f[i]
	rhs := f[j]

	lhsIsDir, rhsIsDir := lhs.IsDir(), rhs.IsDir()
	if lhsIsDir && !rhsIsDir {
		return true
	} else if !lhsIsDir && rhsIsDir {
		return false
	} else {
		return sort.StringSlice([]string{lhs.Name(), rhs.Name()}).Less(0, 1)
	}
}

func (f Files) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
