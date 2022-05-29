package main

import (
	"log"
	"net"

	"itumate.com/im/conf"
	"itumate.com/im/server/pool"
)

func main() {

	network := conf.GetNetwork()
	host := conf.GetHost()

	l, err := net.Listen(network, host)
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

		go pool.ProcessConn(c)
	}

}
