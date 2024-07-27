package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/svbnbyrk/go-ddd/internal/app"
	"github.com/svbnbyrk/go-ddd/internal/app/command"
	"github.com/svbnbyrk/go-ddd/pkg/server"
)

func main() {
	r := chi.NewRouter()

	a := &app.Application{
		Commands: app.Commands{
			CreateWalletHandler: command.NewCreateWalletHandler(),
		},
		Queries: app.Queries{},
	}

	r.Post("/wallets", server.Serve[command.CreateWalletRequest, command.CreateWalletResponse](a.Commands.CreateWalletHandler))

	srv := server.New(r)

	go func() {
		if err := <-srv.Notify(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := srv.Shutdown(); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
