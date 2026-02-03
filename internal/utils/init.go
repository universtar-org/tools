package utils

import (
	"context"

	"github.com/universtar-org/tools/internal/api"
)

func InitClientAndContext(token string) (*api.Client, context.Context) {
	client := api.NewClient(token)
	ctx := context.Background()

	return client, ctx
}
