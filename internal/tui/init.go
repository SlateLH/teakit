package tui

import (
	"fmt"
	"path/filepath"

	"github.com/SlateLH/teakit/internal/commands"
	"github.com/SlateLH/teakit/internal/config"
	"github.com/charmbracelet/huh"
)

var (
	form          *huh.Form
	componentsDir string
)

func init() {
	form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Components directory").
				Value(&componentsDir).
				Placeholder(filepath.Join(config.DefaultTUIDir, config.DefaultComponentsDir)),
		),
	)
}

func RunInit() error {
	if config.Exists() {
		return fmt.Errorf("teakit is already initialized (found %s)", config.ConfigFile)
	}

	if err := form.Run(); err != nil {
		return err
	}

	if componentsDir == "" {
		componentsDir = filepath.Join(config.DefaultTUIDir, config.DefaultComponentsDir)
	}

	cfg := config.Config{
		ComponentsDir: componentsDir,
		Registries: map[string]string{
			config.TeakitRegistryAlias: config.TeakitRegistry,
		},
	}

	if err := commands.Init(cfg); err != nil {
		return err
	}

	return nil
}
