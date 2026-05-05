// Copyright (c) 2024 Tom Payne
// SPDX-License-Identifier: MIT

// chezmoi manages your dotfiles across multiple machines, securely.
package main

import (
	"fmt"
	"os"

	"github.com/twpayne/chezmoi/v2/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		// Use stderr for error output and exit with a non-zero status code.
		// Exit code 1 is used for general errors; specific exit codes may be
		// returned by subcommands via cmd.ExitCodeError.
		fmt.Fprintf(os.Stderr, "chezmoi: error: %v\n", err)
		os.Exit(1)
	}
}
