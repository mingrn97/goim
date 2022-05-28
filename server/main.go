package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"itumate.com/im/config"
	"log"
	"net"
)

func main() {

	l, err := net.Listen(config.Network, config.Host)
	if err != nil {
		log.Fatalln("start server fail", err)
	}

	defer func(l net.Listener) {
		_ = l.Close()
	}(l)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept connect fail", err)
			continue
		}

		go ProcessConn(c)
	}

}

func ProcessConn(c net.Conn) {
	defer func(c net.Conn) {
		_ = c.Close()
	}(c)

	addr := c.RemoteAddr()
	log.Println("remote client", addr.Network(), addr.String(), "establish connect!")

	buf := make([]byte, 10)
	for {
		_, err := c.Read(buf)
		if err == io.EOF {
			log.Println("remote client", addr.Network(), addr.String(), "exit!")
			return
		}
		if err != nil {
			log.Println("read remote client", addr.Network(), addr.String(), "data err!")
		}

		// uint32
		// client
		var rp uint32 = 0
		packetN := binary.BigEndian.Uint32(buf[:4])

		for {
			n, err := c.Read(buf)
			if err == io.EOF {
				log.Println("remote client", addr.Network(), addr.String(), "exit!")
				return
			}
			if err != nil {
				log.Println("read remote client", addr.Network(), addr.String(), "data err!")
			}

			rp += uint32(n)
			if rp >= packetN {
				break
			}

			fmt.Println(string(buf[:n]))
		}

	}
}
