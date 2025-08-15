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
	httpSrv := server.StartHttpServer(":8080")
	tcpCloser := server.StartTcpListener(":9090")
	grpcSrv, grpcLis := server.StartGrpcServer(":50051")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Println("shutting down...")

	// Graceful shutdown
	server.ShutdownHttp(httpSrv)
	server.ShutdownTcp(tcpCloser)
	server.ShutdownGrpc(grpcSrv, grpcLis)
}
