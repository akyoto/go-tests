package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aerogo/nano"
)

var hosts = []string{}

// User ...
type User struct {
	ID        string
	Name      string
	BirthYear string
	Text      string
	Created   string
	Edited    string
	Following []string
}

func main() {
	node := nano.New(5000, hosts...)
	node.Namespace("test", (*User)(nil))
	defer node.Close()

	fmt.Println("Running...")

	i := 0

	for {
		time.Sleep(1000 * time.Millisecond)
		key := "hostname:" + strconv.Itoa(i)
		fmt.Println(key)
		node.Namespace("test").Set("User", key, &User{ID: strconv.Itoa(i)})
		i++
	}

	//stop := make(chan os.Signal)
	//signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	//<-stop
}
