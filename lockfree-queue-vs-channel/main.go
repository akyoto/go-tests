package main

import (
	"fmt"
	"time"

	lfc "github.com/purpuregecko/go-lfc"
	"github.com/scryner/lfreequeue"
)

const n = 100000
const p = 8

func main() {
	for {
		lfcBench()
		lfreequeueBench()
		channelBench()
	}
}

func noop(v interface{}) {}

func lfcBench() {
	q := lfc.NewQueue()
	done := make(chan bool)
	start := time.Now()

	for h := 0; h < p; h++ {
		go func() {
			for i := 0; i < n; i++ {
				q.Enqueue(i)
			}
			done <- true
		}()

		go func() {
			i := 0

			for i < n {
				v, ok := q.Dequeue()

				if !ok {
					continue
				}

				noop(v)
				i++
			}

			done <- true
		}()
	}

	for h := 0; h < p; h++ {
		<-done
		<-done
	}

	fmt.Println(time.Since(start), "queue (lfc)")
}

func lfreequeueBench() {
	q := lfreequeue.NewQueue()
	done := make(chan bool)
	start := time.Now()

	for h := 0; h < p; h++ {
		go func() {
			for i := 0; i < n; i++ {
				q.Enqueue(i)
			}
			done <- true
		}()

		go func() {
			i := 0

			for i < n {
				v, ok := q.Dequeue()

				if !ok {
					continue
				}

				noop(v)
				i++
			}

			done <- true
		}()
	}

	for h := 0; h < p; h++ {
		<-done
		<-done
	}

	fmt.Println(time.Since(start), "queue (lfreequeue)")
}

func channelBench() {
	q := make(chan int)
	done := make(chan bool)
	start := time.Now()

	for h := 0; h < p; h++ {
		go func() {
			for i := 0; i < n; i++ {
				q <- i
			}
			done <- true
		}()

		go func() {
			for i := 0; i < n; i++ {
				v := <-q
				noop(v)
			}
			done <- true
		}()
	}

	for h := 0; h < p; h++ {
		<-done
		<-done
	}

	fmt.Println(time.Since(start), "channel")
}
