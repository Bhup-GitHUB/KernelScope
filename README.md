# KernelScope

KernelScope is a self-hosted observability project focused on learning Linux systems, Go services, and production-style infrastructure. The long-term goal is zero-instrumentation telemetry using eBPF. Phase 1 only contains the repository foundation and local development stack.

## Status

Phase 1 sets up the first working repository baseline:

- a minimal `agent` command
- a minimal `collector` command
- shared version metadata
- local infrastructure with Docker Compose
- CI, linting, and basic project tooling

This phase does not implement eBPF, gRPC, storage integration, dashboards, or production deployment logic.

## Requirements

- Go 1.22+
- Docker with Docker Compose
- `golangci-lint` for `make lint`

## Local Development

```bash
make build
make test
make lint
make dev
go run ./cmd/agent
go run ./cmd/collector
```

The collector listens on `:9090` by default.

Override the listen address with:

```bash
KERNELSCOPE_COLLECTOR_ADDR=:9191 go run ./cmd/collector
```

## Local Infrastructure

`make dev` starts these local dependencies:

- Postgres
- ClickHouse
- Redis

The Compose file lives at `deploy/dev/compose.yaml`.

## Not Implemented Yet

Phase 1 intentionally does not include:

- eBPF programs
- C sources
- gRPC or protobuf
- database schemas or migrations
- Redis integration in Go code
- Kubernetes, Helm, or Terraform
- API or dashboard services
- authentication or config file parsing

## License

KernelScope is licensed under the MIT License. See [LICENSE](LICENSE).
