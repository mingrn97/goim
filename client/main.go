package main

import (
	"log"
	"net"

	"itumate.com/im/client/pipe"
	"itumate.com/im/colorized"
	"itumate.com/im/conf"
)

func main() {

	// establish connect
	c, err := net.Dial(conf.GetNetwork(), conf.GetHost())
	if err != nil {
		log.Fatalln(colorized.RedDark, "Failed to connect to the server", conf.GetNetwork(), conf.GetHost(), colorized.Reset)
	}

	// don't forget close c
	defer func(c net.Conn) {
		_ = c.Close()
	}(c)

	// Create an empty chan to prevent the main thread from shutting down
	var signal = make(chan struct{})

	go pipe.Reader(c, signal)
	go pipe.Writer(c)

	<-signal
	log.Println(colorized.RedDark+colorized.Bold, "The server has been closed, Automatically exits the client program!"+colorized.Reset)
}
