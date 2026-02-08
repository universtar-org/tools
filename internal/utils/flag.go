package utils

import (
	"flag"

	"github.com/universtar-org/tools/internal/model"
)

func ParseFlags() model.Options {
	var opt model.Options
	flag.BoolVar(&opt.Debug, "verbose", false, "enable debug logging")
	flag.BoolVar(&opt.Debug, "v", false, "enable debug logging (shorthand)")
	flag.StringVar(&opt.Token, "t", "", "pass GitHub token")

	flag.Parse()
	return opt
}
