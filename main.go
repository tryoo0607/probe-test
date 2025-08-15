package main

import (
	"context"
	"log"
	"os/signal"
	"probe-test/server"
	"syscall"
)

func main() {
	// 초기 상태값 준비
	server.InitHealthState()

	// 서버 실행
	httpSrv := server.StartHTTPServer()
	tcpCloser := server.StartTCPListener()
	grpcSrv, grpcLis := server.StartGRPCServer()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Println("shutting down...")

	// Graceful shutdown
	server.ShutdownHTTP(httpSrv)
	server.ShutdownTCP(tcpCloser)
	server.ShutdownGRPC(grpcSrv, grpcLis)
}
