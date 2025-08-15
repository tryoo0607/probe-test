package server

import (
	"log"
	"net"
)

type tcpCloser interface{ Close() error }

func StartTcpListener(addr string) tcpCloser {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP", addr)
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				continue
			}
			_ = c.Close() // 열려있기만 하면 probe 통과
		}
	}()
	return ln
}

func ShutdownTcp(c tcpCloser) { _ = c.Close() }
