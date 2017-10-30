package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	filemutex "github.com/alexflint/go-filemutex"
)

const (
	fileName = "data.txt"
	dataN    = 100000
)

var (
	m    *filemutex.FileMutex
	err  error
	data = []byte(strings.Repeat("Lorem ipsum\n", dataN))
)

func main() {
	m, err = filemutex.New("/tmp/test.lock")

	if err != nil {
		log.Fatalln("Directory did not exist or file could not created")
	}

	go write()

	for {
		// Is the data correct on every read access?
		m.RLock()
		readData, _ := ioutil.ReadFile(fileName)
		m.RUnlock()

		time.Sleep(10 * time.Millisecond)

		parts := bytes.Split(readData, []byte{'\n'})

		if len(parts) != dataN+1 {
			fmt.Printf("Lines %d should be %d\n", len(parts), dataN+1)
		}
	}
}

func write() {
	for {
		m.Lock()
		ioutil.WriteFile(fileName, data, 0644)
		m.Unlock()

		time.Sleep(10 * time.Millisecond)
	}
}
