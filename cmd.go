package main

import (
	"fmt"
	"github.com/bbortt/golang-rfc-865/v2/internal"
	"math/rand"
	"net"
)

func main() {
	fmt.Printf("Starting \"Quote Of The Day Service\"..\n")

	l, err := net.Listen("tcp", ":17")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Printf("Server is ready: Awaiting connections..\n")

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	ra := c.RemoteAddr().String()
	fmt.Printf("New connection from %s\n", ra)

	_, err := c.Write([]byte(quoteOfTheDay()))
	if err != nil {
		fmt.Printf("Failed to serve %s:\n\t%s\n", ra, err.Error())
	}

	if err := c.Close(); err != nil {
		fmt.Printf("Connection to %s not closed correctly\n", ra)
	}
}

func quoteOfTheDay() string {
	return internal.Quotes[rand.Intn(len(internal.Quotes))]
}
