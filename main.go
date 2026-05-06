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
		os.Exit(exitCode)
	}
}
