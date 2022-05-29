package pool

import (
	"log"
	"net"

	"itumate.com/im/exp"
	"itumate.com/im/transport"
)

func ProcessConn(c net.Conn) {

	defer func(c net.Conn) {
		_ = c.Close()
	}(c)

	clientLogin(c)

	addr := c.RemoteAddr()
	log.Println("remote client", addr.Network(), addr.String(), "establish connect!")

	for {
		message, err := transport.Read(c)
		if err == exp.ConnClose {
			log.Println("remote client", addr.Network(), addr.String(), "exit!")
			clientLogout(c)
			return
		}
		if err != nil {
			log.Println("read remote client", addr.Network(), addr.String(), "data exp!")
			continue
		}

		log.Printf("%#v", message)

		broadcast(c, message)

	}
}
