package server

import (
	"log"
	"net"
	"probe-test/config"
	"probe-test/util"

	"google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func StartGRPCServer() (*grpc.Server, net.Listener) {
	// 설정 값 불러오기
	cfg := config.GetInstance()
	port := util.ConvertToPortString(cfg.TCPPort)

	lis, err := net.Listen("tcp", port)
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

			if ready.Load() {
				hs.SetServingStatus("ready", healthpb.HealthCheckResponse_SERVING)
			} else {
				hs.SetServingStatus("ready", healthpb.HealthCheckResponse_NOT_SERVING)
			}

			if started.Load() {
				hs.SetServingStatus("startup", healthpb.HealthCheckResponse_SERVING)
			} else {
				hs.SetServingStatus("startup", healthpb.HealthCheckResponse_NOT_SERVING)
			}

			// 너무 자주 돌지 않게 간단히 sleep
			util.SleepMs(500)
		}
	}()

	go func() {
		log.Println("gRPC", port)
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	return s, lis
}

func ShutdownGRPC(s *grpc.Server, lis net.Listener) {
	s.GracefulStop()
	_ = lis.Close()
}
