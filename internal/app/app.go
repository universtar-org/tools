package app

import (
	"context"

	"github.com/universtar-org/tools/internal/api"
)

type App struct {
	Client *api.Client
	Ctx    context.Context
}

func New(token string) *App {
	return &App{
		Client: api.NewClient(token),
		Ctx:    context.Background(),
	}
}
