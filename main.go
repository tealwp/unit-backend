package main

import (
	"github.com/tealwp/unit-backend/cmd"
	"github.com/tealwp/unit-backend/pkg/cfg"
)

func main() {
	cfg := cfg.NewConfig()
	if err := cmd.App(cfg); err != nil {
		panic(err)
	}
}
