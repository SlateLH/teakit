package commands

import (
	"fmt"
	"os"

	"github.com/SlateLH/teakit/internal/config"
)

func Init(cfg config.Config) error {
	if config.Exists() {
		return fmt.Errorf("teakit is already initialized (found %s)", config.ConfigFile)
	}

	if err := os.MkdirAll(cfg.ComponentsDir, 0755); err != nil {
		return err
	}

	if err := config.Write(config.ConfigFile, cfg); err != nil {
		return err
	}

	return nil
}
