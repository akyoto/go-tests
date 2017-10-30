package main

import "time"

var channel = make(chan bool)

func main() {
	go consume()
	time.Sleep(500 * time.Millisecond)
	close(channel)
	time.Sleep(500 * time.Millisecond)
}

func consume() {
	for _ = range channel {
		println("consumed")
	}

	println("closed")
}
