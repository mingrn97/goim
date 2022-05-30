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

		// Send the read data to the server
		transport.WritePlaintext(c, line)
	}
}
