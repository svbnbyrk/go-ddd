package command

import (
	"context"
	"net/http"

	"github.com/svbnbyrk/go-ddd/pkg/server"
)

type CreateWalletRequest struct {
	WalletName string
}
type CreateWalletResponse struct{}
type CreateWalletHandler struct{}

func NewCreateWalletHandler() *CreateWalletHandler {
	return &CreateWalletHandler{}
}

func (h *CreateWalletHandler) Handle(ctx context.Context, req *CreateWalletRequest) (*CreateWalletResponse, error) {
	if req.WalletName == "" {
		return nil, server.NewAppError(http.StatusBadRequest, "Wallet is required")
	}

	return &CreateWalletResponse{}, nil
}
