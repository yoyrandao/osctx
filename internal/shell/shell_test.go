package shell

import "testing"

func TestExportStmt_POSIX(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	t.Setenv("PSModulePath", "")
	got := ExportStmt("dev")
	if got != "export OS_CLOUD=dev" {
		t.Errorf("got %q, want %q", got, "export OS_CLOUD=dev")
	}
}

func TestExportStmt_PowerShell(t *testing.T) {
	t.Setenv("SHELL", "")
	t.Setenv("PSModulePath", "/usr/local/share/powershell/Modules")
	got := ExportStmt("dev")
	want := `$env:OS_CLOUD = "dev"`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUnsetStmt_POSIX(t *testing.T) {
	t.Setenv("SHELL", "/bin/zsh")
	t.Setenv("PSModulePath", "")
	got := UnsetStmt()
	if got != "unset OS_CLOUD" {
		t.Errorf("got %q, want %q", got, "unset OS_CLOUD")
	}
}

func TestUnsetStmt_PowerShell(t *testing.T) {
	t.Setenv("SHELL", "")
	t.Setenv("PSModulePath", "/usr/local/share/powershell/Modules")
	got := UnsetStmt()
	want := `Remove-Item Env:\OS_CLOUD`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
