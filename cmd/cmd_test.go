package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// setupClouds writes a minimal clouds.yaml in dir and chdirs into it.
func setupClouds(t *testing.T, names ...string) {
	t.Helper()
	dir := t.TempDir()
	var sb strings.Builder
	sb.WriteString("clouds:\n")
	for _, n := range names {
		sb.WriteString("  " + n + ":\n    auth: {}\n")
	}
	if err := os.WriteFile(filepath.Join(dir, "clouds.yaml"), []byte(sb.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	orig, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(orig) }) //nolint:errcheck
	os.Chdir(dir)                        //nolint:errcheck
}

func runCmd(args []string) (stdout, stderr string, err error) {
	outBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs(args)
	err = rootCmd.Execute()
	return outBuf.String(), errBuf.String(), err
}

func TestLsCmd_ListsClouds(t *testing.T) {
	setupClouds(t, "dev", "prod")
	_, _, err := runCmd([]string{"ls"})
	if err != nil {
		t.Fatalf("osctx ls: %v", err)
	}
}

func TestCurrentCmd_ReadsOSCloud(t *testing.T) {
	t.Setenv("OS_CLOUD", "dev")
	_, _, err := runCmd([]string{"current"})
	if err != nil {
		t.Fatalf("osctx current: %v", err)
	}
}

func TestCurrentCmd_NoneWhenUnset(t *testing.T) {
	t.Setenv("OS_CLOUD", "")
	_, _, err := runCmd([]string{"current"})
	if err != nil {
		t.Fatalf("osctx current: %v", err)
	}
}

func TestUnsetCmd_WhenCloudSet(t *testing.T) {
	t.Setenv("OS_CLOUD", "dev")
	t.Setenv("SHELL", "/bin/sh")

	// Capture real os.Stdout because unset writes directly there (shell sources it).
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	origStdout := os.Stdout
	os.Stdout = w
	t.Cleanup(func() { os.Stdout = origStdout })

	if _, _, err = runCmd([]string{"unset"}); err != nil {
		t.Fatalf("osctx unset: %v", err)
	}
	w.Close()

	var buf bytes.Buffer
	buf.ReadFrom(r) //nolint:errcheck
	if !strings.Contains(buf.String(), "unset OS_CLOUD") {
		t.Errorf("expected 'unset OS_CLOUD' on stdout, got %q", buf.String())
	}
}

func TestUnsetCmd_WhenCloudNotSet(t *testing.T) {
	t.Setenv("OS_CLOUD", "")
	_, _, err := runCmd([]string{"unset"})
	if err != nil {
		t.Fatalf("osctx unset with no cloud: %v", err)
	}
}
