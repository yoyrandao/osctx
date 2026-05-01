package selector

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/yoyrandao/osctx/internal/utils"
)

// Select presents clouds to the user and returns the chosen cloud name.
// It uses fzf when available and falls back to a numbered-list prompt.
func Select(clouds []string, stderr io.Writer) (string, error) {
	if len(clouds) == 0 {
		return "", fmt.Errorf("no clouds available")
	}
	if hasFzf() {
		return runFzf(clouds)
	}
	return runFallback(clouds, stderr)
}

func hasFzf() bool {
	if os.Getenv("OSCTX_IGNORE_FZF") == "true" {
		return false
	}

	_, err := exec.LookPath("fzf")
	return err == nil
}

func runFzf(clouds []string) (string, error) {
	input := strings.Join(clouds, "\n")
	cmd := exec.Command(
		"fzf",
		"--ansi",
		"--no-multi",
		"--layout=reverse",
		"--height=40%",
		"--layout=reverse",
		"--prompt=OS_CLOUD> ",
		"--header="+fmt.Sprintf("current: %s", utils.GetOSCloud()),
	)
	cmd.Stdin = bytes.NewBufferString(input)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		// fzf exits with code 130 when the user presses Escape / Ctrl-C.
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 130 {
			return "", fmt.Errorf("selection cancelled")
		}
		return "", fmt.Errorf("fzf: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}

func runFallback(clouds []string, stderr io.Writer) (string, error) {
	for i, name := range clouds {
		fmt.Fprintf(stderr, "%d) %s\n", i+1, name)
	}
	fmt.Fprintf(stderr, "Select cloud [1-%d]: ", len(clouds))

	var raw string
	fmt.Fscan(os.Stdin, &raw)

	idx, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil || idx < 1 || idx > len(clouds) {
		return "", fmt.Errorf("invalid selection %q", raw)
	}

	return clouds[idx-1], nil
}
