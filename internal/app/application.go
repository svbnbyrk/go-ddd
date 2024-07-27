package app

import "github.com/svbnbyrk/go-ddd/internal/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateWalletHandler *command.CreateWalletHandler
}

type Queries struct {
}
