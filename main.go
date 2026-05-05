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
		fmt.Fprintf(os.Stderr, "chezmoi: error: %v\n", err)
		os.Exit(1)
	}
}
