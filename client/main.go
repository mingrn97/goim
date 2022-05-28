package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"itumate.com/im/config"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	c, err := net.Dial(config.Network, config.Host)
	if err != nil {
		log.Fatalln("Failed to connect to the server", config.Network, config.Host)
	}
	defer func(c net.Conn) {
		_ = c.Close()
	}(c)

	addr := c.RemoteAddr()
	log.Println("remote server", addr.Network(), addr.String(), "establish connect!")

	r := bufio.NewReader(os.Stdin)

	for {
		line, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		//line = strings.Trim(line, " \n\r")
		//if line == "exit" {
		//	break
		//}

		// data length
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], uint32(len(line)))

		_, _ = c.Write(buf[:])
		_, _ = c.Write(line)

	}

	time.Sleep(time.Second * 3)

}
