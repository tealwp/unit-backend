package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/tealwp/unit-backend/pkg/cfg"
	"github.com/tealwp/unit-backend/pkg/grpc"
)

func App(cfg *cfg.Config) error {
	flag.Parse()

	if len(flag.Args()) < 1 {
		return errors.New("no arguments for cli")
	}

	command := flag.Args()[0]
	switch command {
	case "grpc":
		err := grpc.Serve(cfg)
		if err != nil {
			return err
		}
	case "migrate":
		err := Migrate(cfg)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cli command: %s is not a proper input", command)
	}
	return nil
}
