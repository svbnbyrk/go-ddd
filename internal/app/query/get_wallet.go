package query

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/svbnbyrk/go-ddd/pkg/server"
)

type GetWalletRequest struct {
	ID uuid.UUID
}
type GetWalletResponse struct{}
type GetWalletHandler struct{}

func NewGetWalletHandler() *GetWalletHandler {
	return &GetWalletHandler{}
}

func (h *GetWalletHandler) Handle(ctx context.Context, req *GetWalletRequest) (*GetWalletResponse, error) {
	if req.ID.String() == "" {
		return nil, server.NewAppError(http.StatusBadRequest, "Wallet is required")
	}

	return &GetWalletResponse{}, nil
}
