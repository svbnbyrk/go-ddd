package command

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/svbnbyrk/go-ddd/pkg/server"
)

type CreateWalletRequest struct {
	WalletName string
}
type CreateWalletResponse struct {
	ID uuid.UUID
}
type CreateWalletHandler struct{}

func NewCreateWalletHandler() *CreateWalletHandler {
	return &CreateWalletHandler{}
}

func (h *CreateWalletHandler) Handle(ctx context.Context, req *CreateWalletRequest) (*CreateWalletResponse, error) {
	if req.WalletName == "" {
		return nil, server.NewAppError(http.StatusBadRequest, "Wallet is required")
	}

	return &CreateWalletResponse{
		ID: uuid.New(),
	}, nil
}
