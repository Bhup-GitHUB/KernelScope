package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bhup-GitHUB/kernelscope/internal/version"
)

func main() {
	addr := os.Getenv("KERNELSCOPE_COLLECTOR_ADDR")
	if addr == "" {
		addr = ":9090"
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := &http.Server{
		Addr:    addr,
		Handler: newMux(),
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)

	go func() {
		logger.Info("collector starting", "addr", addr, "version", version.Version, "commit", version.Commit, "date", version.Date)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		logger.Info("collector shutting down")
	case err := <-errCh:
		if err != nil {
			logger.Error("collector stopped", "error", err)
			os.Exit(1)
		}
		return
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("collector shutdown failed", "error", err)
		os.Exit(1)
	}

	if err := <-errCh; err != nil {
		logger.Error("collector stopped", "error", err)
		os.Exit(1)
	}
}

func okHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func metricsHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("kernelscope metrics placeholder\n"))
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", okHandler)
	mux.HandleFunc("/readyz", okHandler)
	mux.HandleFunc("/metrics", metricsHandler)
	return mux
}
