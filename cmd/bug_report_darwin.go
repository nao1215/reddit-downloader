//go:build darwin

// Package cmd is a package that contains subcommands for the reddit-dl CLI command.
package cmd

import (
	"os/exec"
)

func openBrowser(targetURL string) bool {
	return exec.Command("open", targetURL).Start() == nil
}
