package rest

import (
	"context"
	"file-server/cmd/rest/middleware"
	"file-server/lib/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type handler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	GetFile(w http.ResponseWriter, r *http.Request)
}

func Run(ctx context.Context, cfg *config.Config, h handler) error {
	router := mux.NewRouter()

	authMiddleware := middleware.NewAuthMiddleware(cfg.APIConfig)

	privateAPI := router.PathPrefix("/api/v1").Subrouter()
	privateAPI.Use(authMiddleware.Middleware)
	privateAPI.HandleFunc("/upload", h.UploadFile).Methods("POST")

	publicAPI := router.PathPrefix("/public").Subrouter()
	publicAPI.HandleFunc("/file/storage/{tipe}/{filename}", h.GetFile).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Mode", "Accept-Encoding", "Accept-Language", "Connection", "Host", "Origin", "Referer", "User-Agent", "X-Requested-With"},
		MaxAge:             120, // 1 minutes
		AllowCredentials:   true,
		OptionsPassthrough: false,
		Debug:              false,
	})

	httpHandler := c.Handler(router)

	if err := startServer(ctx, httpHandler, cfg); err != nil {
		return err
	}

	return nil
}

func startServer(ctx context.Context, httpHandler http.Handler, cfg *config.Config) error {
	errChan := make(chan error)
	go func() {
		errChan <- StartHTTP(ctx, httpHandler, cfg)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func StartHTTP(ctx context.Context, httpHandler http.Handler, cfg *config.Config) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.APPConfig.HTTPPort),
		Handler: httpHandler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server started on port %d", cfg.APPConfig.HTTPPort)
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruption

	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server stopped")
	return nil
}
