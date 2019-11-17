package chiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rianekacahya/config"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	custom_middleware "github.com/rianekacahya/chiserver/middleware"
)

var (
	server *chi.Mux
	mutex  sync.Once
)

func GetServer() *chi.Mux {
	mutex.Do(func() {
		server = newServer()
	})
	return server
}

func newServer() *chi.Mux {
	return chi.NewRouter()
}

func InitServer() {

	// init default middleware
	GetServer().Use(middleware.Recoverer)
	GetServer().Use(custom_middleware.CORS)
	GetServer().Use(custom_middleware.Recovery)
	GetServer().Use(custom_middleware.Headers)
	GetServer().Use(custom_middleware.Logger)

	// healthCheck endpoint
	GetServer().Get("/infrastructure/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal("OK")
		w.Write(response)
	})
}

func StartServer(ctx context.Context) {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", config.GetChiServerPort()),
		Handler: GetServer(),
		ReadTimeout:  time.Duration(config.GetHTTPServerReadTimeout()) * time.Second,
		WriteTimeout: time.Duration(config.GetHTTPServerWriteTimeout()) * time.Second,
		IdleTimeout:  time.Duration(config.GetHTTPServerIdleTimeout()) * time.Second,
	}

	select {
	case <-ctx.Done():
		if err := srv.Shutdown(ctx); err != nil {
			panic(err)
		}
	default:
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}
}