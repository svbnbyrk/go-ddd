package server

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	readTimeout     = 5 * time.Second
	writeTimeout    = 5 * time.Second
	addr            = ":8081"
	shutdownTimeout = 5 * time.Second
)

// Server
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// Handler interface
type Handler[I any, O any] interface {
	Handle(context.Context, *I) (*O, error)
}

// Serve
func Serve[I any, O any](path string, h Handler[I, O]) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req I

		if err := c.Bind(&req); err != nil {
			return err
		}

		handle, err := h.Handle(c.Request().Context(), &req)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, handle)
	}
}

// New
func New() *Server {
	e := echo.New()

	httpServer := &http.Server{
		Handler:      e,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Addr:         addr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: shutdownTimeout,
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
