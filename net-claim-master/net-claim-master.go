package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aerogo/packet"
)

const (
	address   = ":5000"
	nodeCount = 2
)

func main() {
	for i := 0; i < nodeCount; i++ {
		go startNode(i)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}

func startNode(id int) {
	listener, err := net.ListenPacket("udp", address)

	if err != nil {
		fmt.Printf("[%d] Is client\n", id)
		client(id)
		return
	}

	fmt.Printf("[%d] Is server\n", id)
	server(id, listener)
}

func client(id int) {
	var (
		connection *net.UDPConn
		udpAddress *net.UDPAddr
		err        error
	)

	for {
		udpAddress, err = net.ResolveUDPAddr("udp", address)

		if err != nil {
			fmt.Printf("[%d] Error resolving address: %v\n", id, err)
			continue
		}

		connection, err = net.DialUDP("udp", nil, udpAddress)

		if err != nil {
			fmt.Printf("[%d] Error connecting to server: %v\n", id, err)
			continue
		}

		break
	}

	fmt.Printf("[%d] Successfully connected to server %v\n", id, connection.RemoteAddr())
	defer connection.Close()

	stream := packet.NewStream(0)

	stream.OnError(func(ioErr packet.IOError) {
		fmt.Printf("[%d] Error sending message to server: %v\n", id, ioErr.Error)
	})

	stream.SetConnection(connection)
	stream.Outgoing <- packet.New(0, []byte("ping"))
	time.Sleep(time.Hour)

	// _, err = connection.Write([]byte("ping"))

	// if err != nil {
	// 	fmt.Printf("[%d] Error sending message to server: %v\n", id, err)
	// }
}

func server(id int, listener net.PacketConn) {
	buffer := make([]byte, 4096)

	for {
		n, address, err := listener.ReadFrom(buffer)

		if err != nil {
			fmt.Printf("[%d] Error reading client packet: %v\n", id, err)
			continue
		}

		fmt.Printf("[%d] Read %d bytes from %v\n", id, n, address)
	}
}
