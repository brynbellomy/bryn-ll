package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	x, _ := ioutil.ReadDir(os.Getenv("HOME"))
	for _, file := range x {
		fmt.Printf("%s [%v] %v\n", file.Name(), file.Size(), file.Mode())
	}
}
