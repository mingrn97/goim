package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"itumate.com/im/conf"
	"itumate.com/im/exp"
	"itumate.com/im/transport"
)

func main() {

	// establish connect
	c, err := net.Dial(conf.GetNetwork(), conf.GetHost())
	if err != nil {
		log.Fatalln("Failed to connect to the server", conf.GetNetwork(), conf.GetHost())
	}

	// don't forget close c
	defer func(c net.Conn) {
		_ = c.Close()
	}(c)

	// Create a empty chan to prevent the main thread from shutting down
	var signal = make(chan struct{})

	go reader(c)
	go writer(c)

	<-signal
}

func reader(c net.Conn) {

	addr := c.RemoteAddr()
	log.Println("remote server", addr.Network(), addr.String(), "establish connect!")

	for {
		// Read data from server.
		// If the data is read, the data is encapsulated into a packet(transport.Message)
		packet, err := transport.Read(c)
		if err == exp.ConnClose {
			log.Println("remote server", addr.Network(), addr.String(), "Closed!")
			return
		}
		if err != nil {
			log.Println("read remote server", addr.Network(), addr.String(), "data exp!")
			continue
		}

		// print data
		log.Printf("%#v\n", packet)
	}
}

func writer(c net.Conn) {

	// Read message from stdio(Usually console)
	r := bufio.NewReader(os.Stdin)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		line = strings.Trim(line, "\n\r")

		if line == "" {
			log.Println("Please enter some text!")
			continue
		}

		// Send the read data to the server
		transport.WritePlaintext(c, line)
	}
}
