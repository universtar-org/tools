package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/universtar-org/tools/internal/model"
)

// GetRepo Get repo information including description, number of stars, etc., via GitHub API.
func (c *Client) GetRepo(ctx context.Context, owner, repo string) (*model.Repo, int, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/repos/%s/%s", owner, repo))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var r model.Repo
	status, err := c.do(req, &r)
	if err != nil {
		return nil, status, err
	}

	return &r, status, nil
}

func (c *Client) GetRepoByUser(ctx context.Context, username string) ([]model.Repo, int, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/users/%s/repos", username))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var repos []model.Repo
	status, err := c.do(req, &repos)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return repos, status, nil
}

func (c *Client) GetDirContent(ctx context.Context, username, repo, path string) ([]string, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/repos/%s/%s/contents/%s", username, repo, path))
	if err != nil {
		return nil, err
	}

	var contents []map[string]any
	var result []string
	status, err := c.do(req, &contents)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("Get: %v", status)
	}

	for _, v := range contents {
		result = append(result, v["name"].(string))
	}

	return result, nil
}
