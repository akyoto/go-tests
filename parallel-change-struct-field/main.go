package main

import (
	"runtime"
)

// User ...
type User struct {
	Name   string
	Number int64
}

var names = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
var numbers = []int64{1, 2, 3, 4, 5, 6, 7, 8}
var user = &User{
	Name:   names[0],
	Number: numbers[0],
}

func main() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go change(i)
	}

	for {
		current := user.Name
		ok := false

		for _, val := range names {
			if current == val {
				ok = true
				break
			}
		}

		if !ok {
			panic(current)
		}
	}
}

func change(id int) {
	for {
		user.Name = names[id%len(names)]
	}
}
