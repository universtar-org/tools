package io

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/universtar-org/universtar-updater/internal/model"
)

func WriteYaml(projects []model.Project, path string) error {
	data, err := yaml.Marshal(projects)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
