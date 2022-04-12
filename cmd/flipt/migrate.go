package main

import (
	"errors"
	"fmt"

	"github.com/markphelps/flipt/config"
	"github.com/markphelps/flipt/storage/sql"
)

var forceMigrate bool

func runMigrate(args []string) error {
	if mode == config.ReadOnlyMode {
		return errors.New("migrate is disabled in read-only mode")
	}

	migrator, err := sql.NewMigrator(*cfg, l)
	if err != nil {
		return fmt.Errorf("creating migrator: %w", err)
	}

	defer migrator.Close()

	return migrator.Run(true)
}
