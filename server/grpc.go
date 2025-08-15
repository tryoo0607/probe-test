package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func StartGrpcServer(addr string) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	// Health 서비스 등록 (k8s gRPC probe 대응)
	hs := grpchealth.NewServer()
	healthpb.RegisterHealthServer(s, hs)

	// liveness 토글에 맞춰 상태 바꾸고 싶다면:
	go func() {
		for {
			if alive.Load() {
				hs.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
			} else {
				hs.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
			}
			// 너무 자주 돌지 않게 간단히 sleep
			sleepMs(500)
		}
	}()

	go func() {
		log.Println("gRPC", addr)
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	return s, lis
}

func ShutdownGrpc(s *grpc.Server, lis net.Listener) {
	s.GracefulStop()
	_ = lis.Close()
}
