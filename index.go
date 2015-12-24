package main

import "os"

func main() {
	dir, err := getDir()
	if err != nil {
		panic(err)
	}

	el := &Ellel{
		Dir:       dir,
		Formatter: &DefaultRowFormatter{},
		Theme:     &DefaultTheme{},
		Output:    &DefaultOutput{},
	}

	el.Render()
}

func getDir() (string, error) {
	if len(os.Args) == 1 {
		return os.Getwd()
	} else {
		return os.Args[1], nil
	}
}
