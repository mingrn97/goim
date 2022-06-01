package pipe

import (
	"bufio"
	"io"
	"log"
	"net"

	"itumate.com/im/colorized"
	"itumate.com/im/transport"
)

func Reader(c net.Conn, signal chan struct{}) {

	addr := c.RemoteAddr()
	log.Println(colorized.GreenDark, "Remote server", addr.Network(), addr.String(), "establish connect!", colorized.Reset)

	r := bufio.NewReader(c)
	for {
		message, err := transport.Decode(r)
		if err == io.EOF {
			log.Println(colorized.RedDark, "Remote server", addr.Network(), addr.String(), "Closed!", colorized.Reset)
			close(signal)
			break
		}
		if err != nil {
			log.Println("read remote server", addr.Network(), addr.String(), "data exp!")
			continue
		}
		// print data
		log.Printf("%#v\n", message)
	}
}
