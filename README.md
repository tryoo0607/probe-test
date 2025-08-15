# Kubernetes Pod Probe 테스트용 어플리케이션

## 개요
이 프로젝트는 Kubernetes 환경에서 **liveness**, **readiness**, **startup** 프로브 동작을 테스트하기 위한 **다중 서버 예제**입니다.  
HTTP, TCP, gRPC 서버를 동시에 실행하며, 각 프로브 상태를 **환경변수로 지정한 지연 시간** 이후 `true`로 전환합니다.

### 주요 기능
- **HTTP 서버**
  - `/healthz` — liveness probe
  - `/readyz` — readiness probe
  - `/startupz` — startup probe
- **TCP 리스너**
  - 단순 연결 후 즉시 닫음
  - TCP probe 통과 여부 확인
- **gRPC 서버**
  - `grpc.health.v1` Health 서비스 구현
  - 프로브별 서비스명 지정 가능 (`LIVENESS_SERVICE_NAME`, `READINESS_SERVICE_NAME`, `STARTUP_SERVICE_NAME`)
- **지연 설정**
  - 각 프로브 활성화 시점을 초 단위 환경변수로 설정 가능
  - 공통 지연(`PROBE_DELAY_SEC`)과 개별 지연 동시 지원
- **Graceful Shutdown**
  - SIGINT, SIGTERM 수신 시 HTTP, TCP, gRPC 서버 순차 종료

## 환경 변수 (ENV)

| Name                        | Type         | Default           | Required | Description                                                                 |
| --------------------------- | ------------ | ----------------- | :------: | --------------------------------------------------------------------------- |
| `HTTP_PORT`                 | string(port) | `8080`            | ☐        | HTTP 서버 포트 예: `"8080"`                                                  |
| `TCP_PORT`                  | string(port) | `9090`            | ☐        | TCP 리스너 포트 예: `"9090"`                                                 |
| `GRPC_PORT`                 | string(port) | `50051`           | ☐        | gRPC 서버 포트 예: `"50051"`                                                 |
| `LIVENESS_SERVICE_NAME`     | string       | `""`              | ☐        | gRPC Health에서 **전체 상태**를 나타내는 서비스명. 빈 문자열 `""`이 표준(권장) |
| `READINESS_SERVICE_NAME`    | string       | `ready`           | ☐        | gRPC Health에서 readiness 상태 서비스명                                     |
| `STARTUP_SERVICE_NAME`      | string       | `startup`         | ☐        | gRPC Health에서 startup 상태 서비스명                                       |
| `PROBE_DELAY_SEC`           | int(seconds) | `0`               | ☐        | 세 프로브 공통 지연(초). 개별 값이 없을 때 기본으로 사용                     |
| `PROBE_DELAY_LIVENESS_SEC`  | int(seconds) | `PROBE_DELAY_SEC` | ☐        | liveness 전용 지연(초)                                                       |
| `PROBE_DELAY_READINESS_SEC` | int(seconds) | `PROBE_DELAY_SEC` | ☐        | readiness 전용 지연(초)                                                      |
| `PROBE_DELAY_STARTUP_SEC`   | int(seconds) | `PROBE_DELAY_SEC` | ☐        | startup 전용 지연(초)                                                        |
