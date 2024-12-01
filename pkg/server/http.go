package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Server struct
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// Handler interface
type Handler[I any, O any] interface {
	Handle(context.Context, *I) (*O, error)
}

func Serve[I any, O any](h Handler[I, O]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req I
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			WriteJSON(w, http.StatusBadRequest, NewAppError(http.StatusBadRequest, err.Error()))
			return
		}

		handle, err := h.Handle(r.Context(), &req)
		if err != nil {
			if appErr, ok := err.(*AppError); ok {
				WriteJSON(w, appErr.Code, appErr)
			} else {
				WriteJSON(w, http.StatusInternalServerError, NewAppError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)))
			}
			return
		}

		WriteJSON(w, http.StatusOK, handle)
	}
}

func New(handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         ":8081",
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: 5 * time.Second,
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

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *AppError:
					http.Error(w, e.Message, e.Code)
				default:
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
