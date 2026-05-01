package clouds

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "clouds.yaml")
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestList_ParsesCloudNames(t *testing.T) {
	yaml := `
clouds:
  dev:
    auth:
      auth_url: https://dev.example.com
  prod:
    auth:
      auth_url: https://prod.example.com
`
	p := writeTemp(t, yaml)

	// Temporarily change working directory so "clouds.yaml" in cwd is found.
	orig, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(orig) }) //nolint:errcheck
	os.Chdir(filepath.Dir(p))            //nolint:errcheck
	// Rename file to match cwd search path name.
	os.Rename(p, filepath.Join(filepath.Dir(p), "clouds.yaml")) //nolint:errcheck

	names, err := List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	sort.Strings(names)
	want := []string{"dev", "prod"}
	if len(names) != len(want) {
		t.Fatalf("got %v, want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Errorf("names[%d] = %q, want %q", i, names[i], want[i])
		}
	}
}

func TestList_NoCloudsFile(t *testing.T) {
	// Point XDG and HOME to a directory with no clouds.yaml, and chdir there.
	dir := t.TempDir()
	orig, _ := os.Getwd()
	t.Cleanup(func() {
		os.Chdir(orig)               //nolint:errcheck
		os.Unsetenv("XDG_CONFIG_HOME") //nolint:errcheck
		os.Unsetenv("HOME")            //nolint:errcheck
	})
	os.Chdir(dir)                          //nolint:errcheck
	os.Setenv("XDG_CONFIG_HOME", dir)      //nolint:errcheck
	os.Setenv("HOME", dir)                 //nolint:errcheck

	_, err := List()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestList_NoCloudsKey(t *testing.T) {
	// Valid YAML but missing the "clouds" key — should return empty list.
	p := writeTemp(t, "other_key: value\n")
	orig, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(orig) }) //nolint:errcheck
	os.Chdir(filepath.Dir(p))            //nolint:errcheck
	os.Rename(p, filepath.Join(filepath.Dir(p), "clouds.yaml")) //nolint:errcheck

	names, err := List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("expected empty names, got %v", names)
	}
}
