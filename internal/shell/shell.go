package shell

import (
	"fmt"
	"os"
	"runtime"
)

type Shell int

const (
	POSIX Shell = iota
	PowerShell
	CMD
)

func Detect() Shell {
	if os.Getenv("SHELL") != "" {
		return POSIX
	}
	if os.Getenv("PSModulePath") != "" {
		return PowerShell
	}
	if runtime.GOOS == "windows" {
		return CMD
	}
	return POSIX
}

func ExportStmt(name string) string {
	switch Detect() {
	case PowerShell:
		return fmt.Sprintf(`$env:OS_CLOUD = "%s"`, name)
	case CMD:
		return fmt.Sprintf("set OS_CLOUD=%s", name)
	default:
		return fmt.Sprintf("export OS_CLOUD=%s", name)
	}
}

func UnsetStmt() string {
	switch Detect() {
	case PowerShell:
		return `Remove-Item Env:\OS_CLOUD`
	case CMD:
		return "set OS_CLOUD="
	default:
		return "unset OS_CLOUD"
	}
}
