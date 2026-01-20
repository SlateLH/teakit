package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigFile           = "teakit.json"
	DefaultTUIDir        = "tui"
	DefaultComponentsDir = "components"
	TeakitRegistry       = "https://github.com/SlateLH/teakit-registry/main"
	TeakitRegistryAlias  = "teakit"
)

type Config struct {
	ComponentsDir string            `json:"componentsDir"`
	Registries    map[string]string `json:"registries"`
}

func Default() Config {
	return Config{
		ComponentsDir: filepath.Join(DefaultTUIDir, DefaultComponentsDir),
		Registries: map[string]string{
			TeakitRegistryAlias: TeakitRegistry,
		},
	}
}

func Write(path string, cfg Config) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(cfg)
}

func Exists() bool {
	_, err := os.Stat(ConfigFile)
	return err == nil
}

func Load() (*Config, error) {
	cfg := Default()

	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("teakit has not been initialized (did not find %s)", ConfigFile)
		}

		return nil, err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Registries == nil {
		cfg.Registries = make(map[string]string)
	}

	return &cfg, nil
}
