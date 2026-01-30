package updater

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/universtar-org/universtar-updater/internal/api"
	"github.com/universtar-org/universtar-updater/internal/io"
)

// Update a file/user.
func Update(client *api.Client, ctx context.Context, path string) error {
	const MAX_TAG_NUMBER = 5
	base := filepath.Base(path)
	owner := strings.TrimSuffix(base, filepath.Ext(base))

	projects, err := io.ReadYaml(path)
	if err != nil {
		return err
	}

	for i := range projects {
		repo, err := client.GetRepo(ctx, owner, projects[i].Repo)
		if err != nil {
			return err
		}

		tag_list := append([]string{repo.Language}, repo.Tags...)
		if len(tag_list) > MAX_TAG_NUMBER {
			tag_list = tag_list[:MAX_TAG_NUMBER]
		}
		projects[i].Description = repo.Description
		projects[i].Stars = repo.Stars
		projects[i].UpdatedAt = repo.UpdatedAt
		projects[i].Tags = tag_list
	}

	io.WriteYaml(projects, path)

	return nil
}
