package registry

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteFile(dest string, data []byte, force bool) error {
	if _, err := os.Stat(dest); err == nil && !force {
		return fmt.Errorf("file %q already exists (use --force to overwrite)", dest)
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}

	return os.WriteFile(dest, data, 0644)
}
