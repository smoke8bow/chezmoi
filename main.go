// Copyright (c) 2024 Tom Payne
// SPDX-License-Identifier: MIT

// chezmoi manages your dotfiles across multiple machines, securely.
package main

import (
	"os"

	"github.com/twpayne/chezmoi/v2/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
