package pipe

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"itumate.com/im/transport"
)

func Writer(c net.Conn) {

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

		var message = &transport.Message{
			Type: transport.PlainType,
			Data: line,
		}

		if data, err := transport.Encode(message); err != nil {
			log.Println("encode data err!", err)
		} else {
			_, _ = c.Write(data)
		}
	}
}
