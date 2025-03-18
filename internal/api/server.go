package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server представляет HTTP-сервер
type Server struct {
	router *mux.Router
	addr   string
}

// NewServer создаёт новый экземпляр Server
func NewServer(handler *Handler, addr string) *Server {
	router := mux.NewRouter()

	router.Use(RequestIDMiddleware)
	router.Use(LoggingMiddleware)

	router.HandleFunc("/api/comments", handler.AddComment).Methods(http.MethodPost)
	router.HandleFunc("/api/comments/news/{news_id:[0-9]+}", handler.GetComments).Methods(http.MethodGet)
	router.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)

	return &Server{
		router: router,
		addr:   addr,
	}
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:         s.addr,
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Сервер комментариев запущен на %s", s.addr)
	return server.ListenAndServe()
}
