package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/universtar-org/universtar-updater/internal/api"
	"github.com/universtar-org/universtar-updater/internal/updater"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("Usage: updater ${data-file-dir}"))
	}
	client := api.NewClient("")
	ctx := context.Background()

	list, err := getDataFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, path := range list {
		fmt.Println("Processing: ", path)
		if err := updater.Update(client, ctx, path); err != nil {
			panic(err)
		}
	}

	fmt.Println("Finished!")
}

func getDataFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, entry := range entries {
		if !entry.IsDir() {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}

	return paths, nil
}
