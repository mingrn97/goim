package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"itumate.com/im/config"
	"itumate.com/im/transport"
	"log"
	"net"
	"os"
	"strings"
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
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		line = strings.Trim(line, "\n\r")

		if line == "" {
			log.Println("please input something")
			continue
		}

		var message = &transport.Message{
			Type: transport.PlainType,
			Data: line,
		}

		data, _ := json.Marshal(message)
		fmt.Println(string(data))

		var me = new(transport.Message)
		_ = json.Unmarshal(data, me)

		// data length
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))

		if n, e := c.Write(buf[:4]); n != 4 || e != nil {
			log.Println("send content length fail", e)
		} else {
			// If the data message length is successfully sent,
			// continue to send the actual data!
			_, _ = c.Write(data)
		}

	}

}
