package main

import (
	"crypto/tls"

	"github.com/lucas-clemente/quic-go"
)

func main() {
	quic.DialAddr("", &tls.Config{}, nil)
}
