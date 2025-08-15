package server

import (
	"log"
	"net"
	"probe-test/config"
	"probe-test/util"
)

type tcpCloser interface{ Close() error }

func StartTCPListener() tcpCloser {
	// 설정 값 불러오기
	cfg := config.GetInstance()
	port := util.ConvertToPortString(cfg.TCPPort)

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP", port)
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

func ShutdownTCP(c tcpCloser) { _ = c.Close() }
