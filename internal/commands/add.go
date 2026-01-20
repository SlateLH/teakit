package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/SlateLH/teakit/internal/config"
	"github.com/SlateLH/teakit/internal/registry"
)

type AddOptions struct {
	Component string
	Registry  string
	Force     bool
}

func resolveRegistryURL(registries map[string]string, alias string) (string, error) {
	if len(registries) == 0 {
		return "", fmt.Errorf("no registries configured in %s", config.ConfigFile)
	}

	if alias != "" {
		url, ok := registries[alias]
		if !ok {
			return "", fmt.Errorf("registry %q not found in %s", alias, config.ConfigFile)
		}

		return url, nil
	}

	if len(registries) == 1 {
		for _, url := range registries {
			return url, nil
		}
	}

	return "", fmt.Errorf("multiple registries configured; specify one using the --registry flag")
}

func Add(opts AddOptions) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	registryURL, err := resolveRegistryURL(cfg.Registries, opts.Registry)
	if err != nil {
		return err
	}

	fetcher, err := registry.NewFetcher(registryURL)
	if err != nil {
		return err
	}

	regMeta, err := fetcher.FetchRegistryMeta()
	if err != nil {
		return err
	}

	component, ok := regMeta.Components[opts.Component]
	if !ok {
		return fmt.Errorf("component %q not found in registry %q", opts.Component, registryURL)
	}

	data, err := fetcher.FetchComponent(component)
	if err != nil {
		return err
	}

	var regAlias string

	if opts.Registry == "" {
		for alias, url := range cfg.Registries {
			if url == registryURL {
				regAlias = alias
				break
			}
		}
	} else {
		regAlias = opts.Registry
	}

	localPath := filepath.Join(cfg.ComponentsDir, regAlias, filepath.Join(strings.Split(component, "/")...))
	if err := registry.WriteFile(localPath, data, opts.Force); err != nil {
		return err
	}

	fmt.Printf("Added component %q from %q\n", opts.Component, registryURL)
	return nil
}
