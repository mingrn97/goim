package pool

import (
	"net"
	"sync"

	"itumate.com/im/transport"
)

var (
	clients clientConn
)

type clientConn struct {
	mu   sync.Mutex
	cons []net.Conn
}

func clientLogin(c net.Conn) {
	clients.mu.Lock()
	defer clients.mu.Unlock()

	clients.cons = append(clients.cons, c)
}

func clientLogout(c net.Conn) {
	clients.mu.Lock()
	defer clients.mu.Unlock()

	l := len(clients.cons)
	for i, con := range clients.cons {
		if con == c {
			if i == 0 {
				clients.cons = clients.cons[1:]
				return
			} else if l == (i + 1) {
				clients.cons = clients.cons[:i]
			} else {
				clients.cons = append(clients.cons[:i], clients.cons[:(i+1)]...)
			}
		}
	}
}

func broadcast(eliminate net.Conn, message *transport.Message) {

	for _, c := range clients.cons {
		if c != eliminate {
			transport.WritePlaintext(c, message.Data)
		}
	}
}
