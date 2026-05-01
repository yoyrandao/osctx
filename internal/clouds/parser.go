package clouds

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type cloudsFile struct {
	Clouds map[string]any `yaml:"clouds"`
}

// searchPaths returns candidate locations for clouds.yaml in priority order.
func searchPaths() []string {
	paths := []string{"clouds.yaml"}

	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		paths = append(paths, filepath.Join(xdg, "openstack", "clouds.yaml"))
	}

	if home, err := os.UserHomeDir(); err == nil {
		paths = append(paths,
			filepath.Join(home, ".config", "openstack", "clouds.yaml"),
			filepath.Join(home, ".openstack", "clouds.yaml"),
		)
	}

	paths = append(paths, "/etc/openstack/clouds.yaml")
	return paths
}

// List finds clouds.yaml and returns the sorted list of cloud names.
func List() ([]string, error) {
	var data []byte
	var readErr error

	for _, p := range searchPaths() {
		data, readErr = os.ReadFile(p)
		if readErr == nil {
			break
		}
	}

	if readErr != nil {
		return nil, fmt.Errorf("clouds.yaml not found in any standard location")
	}

	var cf cloudsFile
	if err := yaml.Unmarshal(data, &cf); err != nil {
		return nil, fmt.Errorf("parse clouds.yaml: %w", err)
	}

	names := make([]string, 0, len(cf.Clouds))
	for name := range cf.Clouds {
		names = append(names, name)
	}

	return names, nil
}
