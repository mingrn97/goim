package conf

import (
	"os"
)

var (
	network string = "tcp"
	host    string = "0.0.0.0:8888"
)

func init() {

	// im_host env can be: 10086 or 0.0.0.0:10086
	if e, exist := os.LookupEnv("im_host"); exist {
		host = e
	}
}

func GetNetwork() string {
	return network
}

func GetHost() string {
	return host
}
