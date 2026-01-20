package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	gitHubPrefix = "https://github.com/"
)

type RegistryMeta struct {
	Components map[string]string `json:"components"`
}

type Fetcher struct {
	Base string
}

func rawBaseURL(repo string) (string, error) {
	if !strings.HasPrefix(repo, gitHubPrefix) {
		return "", fmt.Errorf("only GitHub registries are supported at this time")
	}

	parts := strings.Split(strings.TrimPrefix(repo, gitHubPrefix), "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid GitHub repository URL %q", repo)
	}

	owner := parts[0]
	name := parts[1]
	branch := parts[2]

	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/refs/heads/%s", owner, name, branch), nil
}

func NewFetcher(repo string) (*Fetcher, error) {
	base, err := rawBaseURL(repo)
	if err != nil {
		return nil, err
	}

	return &Fetcher{Base: base}, nil
}

func (f *Fetcher) FetchRegistryMeta() (*RegistryMeta, error) {
	url := f.Base + "/registry.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch registry.json: %s", resp.Status)
	}

	var meta RegistryMeta
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

func (f *Fetcher) FetchComponent(path string) ([]byte, error) {
	url := f.Base + "/" + path

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch component %q: %s", path, resp.Status)
	}

	return io.ReadAll(resp.Body)
}
