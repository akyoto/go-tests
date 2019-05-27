package main

import (
	"net"
	"time"
)

type Node struct {
	listener   net.Listener
	connection net.Conn
}

func (node *Node) IsMaster() bool {
	return node.listener != nil
}

func (node *Node) Close() {
	if node.listener != nil {
		node.listener.Close()
		node.listener = nil
	}

	if node.connection != nil {
		node.connection.Close()
		node.connection = nil
	}
}

func main() {
	// Init
	const n = 10
	nodes := make([]*Node, n)

	for i := 0; i < n; i++ {
		nodes[i] = spawn()
	}

	// Check
	masterNodeCount := 0

	for i := 0; i < n; i++ {
		if nodes[i].IsMaster() {
			masterNodeCount++
		}
	}

	if masterNodeCount != 1 {
		panic("Master node count")
	}

	// Kill master node
	for i := 0; i < n; i++ {
		if nodes[i].IsMaster() {
			nodes[i].Close()
			break
		}
	}

	// Check again
	masterNodeCount = 0

	for i := 0; i < n; i++ {
		if nodes[i].IsMaster() {
			masterNodeCount++
		}
	}

	if masterNodeCount != 1 {
		panic("Master node count wrong after server death")
	}

	// End
	for i := 0; i < n; i++ {
		nodes[i].Close()
	}
}

func spawn() *Node {
	node := &Node{}
	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		connection, err := net.Dial("tcp", ":5000")

		if err != nil {
			panic(err)
		}

		go func() {
			var err error

			for {
				time.Sleep(1 * time.Millisecond)
				node.listener, err = net.Listen("tcp", ":5000")

				if err == nil {
					return
				}
			}
		}()

		node.connection = connection
		return node
	}

	node.listener = listener
	return node
}
