package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/svbnbyrk/go-ddd/internal"
	"github.com/svbnbyrk/go-ddd/internal/ports"
	"github.com/svbnbyrk/go-ddd/pkg/server/router"
)

func main() {
	ctx := context.Background()
	application := internal.NewApplication(ctx)

	router.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(application),
			router,
		)
	})

	log.Print("Server Exited Properly")
}
