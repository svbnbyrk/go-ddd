package wallet

import "context"

type CreateWalletRequest struct{}
type CreateWalletResponse struct{}
type CreateWalletHandler struct {
	Repo string
}

func NewCreateWalletHandler() *CreateWalletHandler {
	return &CreateWalletHandler{}
}

func (h *CreateWalletHandler) Handle(ctx context.Context, req *CreateWalletRequest) (*CreateWalletResponse, error) {
	return &CreateWalletResponse{}, nil
}
