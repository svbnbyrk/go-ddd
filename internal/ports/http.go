package ports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/svbnbyrk/go-ddd/internal/app"
	"github.com/svbnbyrk/go-ddd/internal/app/command"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// CreateWallet handles POST /v1/wallets
func (s HttpServer) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var wallet Wallet
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response, err := s.app.Commands.CreateWalletHandler.Handle(context.Background(), &command.CreateWalletRequest{
		WalletName: wallet.Name,
	})
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetWalletById handles GET /v1/wallets/{id}
func (s HttpServer) GetWalletById(w http.ResponseWriter, r *http.Request, id string) {
	// Example response
	response := WalletResponse{
		Id:      id,
		Name:    "Example Wallet",
		Balance: 100.0,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
