package io

import (
	"github.com/goccy/go-yaml"
	"github.com/universtar-org/universtar-updater/internal/model"
	"os"
)

func ReadYaml(path string) ([]model.Project, error) {
	var projects []model.Project
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
