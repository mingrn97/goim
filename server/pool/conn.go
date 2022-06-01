package pool

import (
	"bufio"
	"io"
	"log"
	"net"

	"itumate.com/im/transport"
)

func ProcessConn(c net.Conn) {

	defer func(c net.Conn) {
		_ = c.Close()
	}(c)

	clientLogin(c)

	addr := c.RemoteAddr()
	log.Println("remote client", addr.Network(), addr.String(), "establish connect!")

	r := bufio.NewReader(c)
	for {
		message, err := transport.Decode(r)
		if err == io.EOF {
			log.Println("remote client", addr.Network(), addr.String(), "exit!")
			clientLogout(c)
			break
		}
		if err != nil {
			log.Println("read remote client", addr.Network(), addr.String(), "data err!", err)
			break
		}

		log.Printf("%#v\n", message)

		broadcast(c, message)
	}
}
