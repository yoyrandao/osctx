package selector

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// patchStdin replaces os.Stdin with a reader containing s for the duration of the test.
func patchStdin(t *testing.T, s string) {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	orig := os.Stdin
	os.Stdin = r
	t.Cleanup(func() {
		os.Stdin = orig
		r.Close()
	})
	w.WriteString(s) //nolint:errcheck
	w.Close()
}

func TestSelect_EmptyClouds(t *testing.T) {
	_, err := Select([]string{}, &bytes.Buffer{})
	if err == nil {
		t.Fatal("expected error for empty clouds")
	}
}

func TestRunFallback_ValidSelection(t *testing.T) {
	clouds := []string{"dev", "staging", "prod"}
	patchStdin(t, "2\n")
	var stderr bytes.Buffer
	got, err := runFallback(clouds, &stderr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "staging" {
		t.Errorf("got %q, want %q", got, "staging")
	}
	if !strings.Contains(stderr.String(), "staging") {
		t.Errorf("stderr should list cloud names, got: %q", stderr.String())
	}
}

func TestRunFallback_InvalidIndex(t *testing.T) {
	clouds := []string{"dev", "prod"}
	patchStdin(t, "99\n")
	_, err := runFallback(clouds, &bytes.Buffer{})
	if err == nil {
		t.Fatal("expected error for out-of-range index")
	}
}

func TestRunFallback_NonNumericInput(t *testing.T) {
	clouds := []string{"dev", "prod"}
	patchStdin(t, "abc\n")
	_, err := runFallback(clouds, &bytes.Buffer{})
	if err == nil {
		t.Fatal("expected error for non-numeric input")
	}
}

func TestRunFallback_FirstAndLast(t *testing.T) {
	clouds := []string{"alpha", "beta", "gamma"}
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"1\n", "alpha"},
		{"3\n", "gamma"},
	} {
		patchStdin(t, tc.input)
		got, err := runFallback(clouds, &bytes.Buffer{})
		if err != nil {
			t.Fatalf("input %q: unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("input %q: got %q, want %q", tc.input, got, tc.want)
		}
	}
}
