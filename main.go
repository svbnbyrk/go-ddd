package main

import (
	"github.com/svbnbyrk/go-ddd/internal/app/wallet"
	"github.com/svbnbyrk/go-ddd/pkg/server"
)

func main() {
	server.New()
	newCreateWalletHandler := wallet.NewCreateWalletHandler()
	server.Serve[wallet.CreateWalletRequest, wallet.CreateWalletResponse]("/wallets", newCreateWalletHandler)
}
