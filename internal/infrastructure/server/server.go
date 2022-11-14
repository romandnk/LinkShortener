package server

import (
	"context"
	"github.com/PudgeRo/LinkShortener/internal/usecases/app/repos/urlrepo"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
	url *urlrepo.Urls
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Start(url *urlrepo.Urls) {
	s.url = url
	go s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}

