package pipe

import (
	"log"
	"net"

	"itumate.com/im/colorized"
	"itumate.com/im/exp"
	"itumate.com/im/transport"
)

func Reader(c net.Conn, signal chan struct{}) {

	addr := c.RemoteAddr()
	log.Println(colorized.GreenDark, "Remote server", addr.Network(), addr.String(), "establish connect!", colorized.Reset)

	for {
		// Read data from server.
		// If the data is read, the data is encapsulated into a packet(transport.Message)
		packet, err := transport.Read(c)
		if err == exp.ConnClose {
			log.Println(colorized.RedDark, "Remote server", addr.Network(), addr.String(), "Closed!", colorized.Reset)
			close(signal)
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
