package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"itumate.com/im/conf"
	"itumate.com/im/transport"
)

func main() {

	c, err := net.Dial(conf.GetNetwork(), conf.GetHost())
	if err != nil {
		log.Fatalln("Failed to connect to the server", conf.GetNetwork(), conf.GetHost())
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
