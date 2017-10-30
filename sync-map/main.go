package main

import (
	"sync"
)

const n = 10000
const channelBufferSize = n

var data sync.Map
var channel chan interface{}

func main() {
	channel = make(chan interface{}, channelBufferSize)

	go write()
	go consume()
	read()
}

func read() {
	for {
		data.Range(func(key, value interface{}) bool {
			channel <- value
			return true
		})
	}
}

func write() {
	for {
		for i := 0; i < n; i++ {
			data.Store(i, i)
		}
	}
}

func consume() {
	for _ = range channel {
		// ...
	}
}
