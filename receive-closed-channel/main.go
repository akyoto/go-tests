package main

func main() {
	channel := make(chan bool)
	close(channel)
	<-channel
}
