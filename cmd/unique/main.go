package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/universtar-org/tools/internal/api"
	"github.com/universtar-org/tools/internal/log"
	"github.com/universtar-org/tools/internal/model"
	"github.com/universtar-org/tools/internal/utils"
)

func main() {
	opts := utils.ParseFlags()
	log.InitLogger(opts.Debug)

	args := flag.Args()
	if len(args) != 1 {
		slog.Error(
			"invalid arguments",
			"usage", "unique ${username}",
		)
		os.Exit(2)
	}

	client, ctx := utils.InitClientAndContext("")
	username := args[0]

	user, err := checkUsername(client, ctx, username)
	if err != nil {
		slog.Error(
			"check username failed",
			"err", err,
		)
	}

	repos, err := client.GetRepoByUser(ctx, username)
	if err != nil {
		slog.Error(
			"get repo by user failed",
			"user", username,
			"err", err,
		)
		os.Exit(1)
	}

	if err := checkUniqueness(client, ctx, repos, *user); err != nil {
		slog.Error(
			"check uniqueness",
			"user", username,
			"err", err,
		)
		os.Exit(1)
	}

	slog.Info("finished")
}

func checkUsername(client *api.Client, ctx context.Context, username string) (*model.User, error) {
	user, err := client.GetUser(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user %s: %w", username, err)
	}
	if username != user.Name {
		return nil, fmt.Errorf("username mismatch: get %s, expect: %s", username, user.Name)
	}
	return user, nil
}

func checkUniqueness(client *api.Client, ctx context.Context, repos []model.Repo, user model.User) error {
	projectWhiteList := []string{"tools", "www"}

	if user.Type != "User" {
		return nil
	}

	owner := "universtar-org"
	path := "data/projects"
	for _, repo := range repos {
		if slices.Contains(projectWhiteList, repo.Name) {
			continue
		}

		slog.Info(
			"checking",
			"repo", owner+"/"+repo.Name,
		)
		contents, err := client.GetDirContent(ctx, owner, repo.Name, path)
		if err != nil {
			return fmt.Errorf("get dir content %s/%s/%s: %w", owner, repo.Name, path, err)
		}

		for _, content := range contents {
			if user.Name == strings.TrimSuffix(content, filepath.Ext(content)) {

				return fmt.Errorf("duplicated username in %s/%s", owner, repo.Name)
			}
		}
	}

	return nil
}
