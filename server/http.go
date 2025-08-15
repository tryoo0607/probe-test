package server

import (
	"log"
	"net/http"
	"probe-test/config"
	"probe-test/util"
	"time"
)

func StartHTTPServer() *http.Server {
	// 설정 값 불러오기
	cfg := config.GetInstance()
	port := util.ConvertToPortString(cfg.HTTPPort)

	// 멀티 플렉스 생성
	mux := http.NewServeMux()

	// liveness
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if !alive.Load() {
			http.Error(w, "not alive", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	// readiness
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if !ready.Load() {
			http.Error(w, "not ready", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	// startup
	mux.HandleFunc("/startupz", func(w http.ResponseWriter, r *http.Request) {
		if !started.Load() {
			http.Error(w, "starting", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	srv := &http.Server{
		Addr:              port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		log.Println("HTTP", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return srv
}

func ShutdownHTTP(srv *http.Server) {
	ctx, cancel := util.TimeoutCtx(5 * time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
