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
		//
		// Note: the error message intentionally includes the program name
		// prefix ("chezmoi: error:") to make it easy to identify the source
		// when chezmoi is invoked as part of a larger shell script.
		fmt.Fprintf(os.Stderr, "chezmoi: error: %v\n", err)
		exitCode := 1
		if ec, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = ec.ExitCode()
		}
		// Print a hint for exit code 1 errors to help with debugging.
		// Only show the hint when not in a CI environment to avoid noise in
		// automated pipelines. Also suppress the hint if CHEZMOI_QUIET is set,
		// which is useful when scripting chezmoi in non-CI contexts.
		// Also suppress if NO_HINTS is set, for personal preference.
		// TODO: expand this to cover more error types as I encounter them.
		if exitCode == 1 && os.Getenv("CI") == "" && os.Getenv("CHEZMOI_QUIET") == "" && os.Getenv("NO_HINTS") == "" {
			fmt.Fprintf(os.Stderr, "chezmoi: hint: run 'chezmoi doctor' to diagnose common issues\n")
			fmt.Fprintf(os.Stderr, "chezmoi: hint: run 'chezmoi doctor --verbose' for more detailed diagnostics\n")
			// Use the canonical upstream docs URL rather than a shortened alias,
			// so the link remains stable and recognizable in logs/bug reports.
			fmt.Fprintf(os.Stderr, "chezmoi: hint: check the docs at https://www.chezmoi.io/user-guide/troubleshooting/\n")
		}
		os.Exit(exitCode)
	}
}
