package internal

import (
	"context"

	"github.com/svbnbyrk/go-ddd/internal/app"
	"github.com/svbnbyrk/go-ddd/internal/app/command"
	"github.com/svbnbyrk/go-ddd/internal/app/query"
)

func NewApplication(ctx context.Context) app.Application {

	return app.Application{
		Commands: app.Commands{
			CreateWalletHandler: command.NewCreateWalletHandler(),
		},
		Queries: app.Queries{
			GetWalletHandler: query.NewGetWalletHandler(),
		},
	}
}
