package main

import (
	"context"
	"fmt"
	"os"

	"github.com/universtar-org/updater/internal/api"
	"github.com/universtar-org/updater/internal/io"
	"github.com/universtar-org/updater/internal/updater"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("Usage: updater ${data-file-dir}"))
	}
	client := api.NewClient("")
	ctx := context.Background()

	list, err := io.GetDataFiles(os.Args[1])
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
