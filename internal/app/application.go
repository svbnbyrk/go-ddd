package app

import (
	"github.com/svbnbyrk/go-ddd/internal/app/command"
	"github.com/svbnbyrk/go-ddd/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateWalletHandler *command.CreateWalletHandler
}

type Queries struct {
	GetWalletHandler *query.GetWalletHandler
}
