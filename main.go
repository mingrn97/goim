package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"itumate.com/im/conf"
	"itumate.com/im/pwd"
)

func main() {

	pwd.NewPwd()
	os.Exit(0)

	listener, err := net.Listen(conf.Network, conf.Host)
	if err != nil {
		fmt.Printf("server start fail: %v\n", err)
	}
	defer listener.Close()

	fmt.Printf("server %s:%s start succeeded!\n", conf.Network, conf.Host)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("server unable to create network connection: %v\n", err)
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	network := addr.Network()
	ip := addr.String()

	fmt.Printf("Connection with client %s:%s succeeded\n", network, ip)

	var buf []byte = make([]byte, 1024)

	for {
		len, err := conn.Read(buf)

		if err == io.EOF {
			fmt.Printf("Client %s:%s say: bye-bye\n", network, ip)
			break
		}

		if err != nil {
			fmt.Printf("Read client %s:%s data err: %v\n", network, ip, err)
		}

		mes := string(buf[:len])

		fmt.Printf("Client %s:%s say: %s\n", network, ip, mes)
	}

}
