package exp

import "errors"

var (
	ConnClose = errors.New("client close connection")
	ReadErr   = errors.New("read client data packets exp")
)
