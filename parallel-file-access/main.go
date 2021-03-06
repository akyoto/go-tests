package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	fileName = "data.txt"
	dataN    = 100000
)

var data = []byte(strings.Repeat("Lorem ipsum\n", dataN))

func main() {
	go write()

	for {
		// Is the data correct on every read access?
		readData, _ := ioutil.ReadFile(fileName)
		parts := bytes.Split(readData, []byte{'\n'})

		if len(parts) != dataN+1 {
			fmt.Printf("Lines %d should be %d\n", len(parts), dataN+1)
		}
	}
}

func write() {
	for {
		ioutil.WriteFile(fileName, data, 0644)
	}
}
