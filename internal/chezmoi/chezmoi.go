// Package chezmoi contains the core logic for chezmoi, a dotfile manager.
package chezmoi

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// Version information, set at build time via ldflags.
var (
	Version   = "dev"
	Commit    = "none"
	Date      = "unknown"
	BuiltBy   = "unknown"
)

// DefaultSourceDirName is the default name of the source directory within the
// user's home config directory.
const DefaultSourceDirName = "chezmoi"

// DefaultDestDirName is the default destination directory (the user's home).
const DefaultDestDirName = "~"

// DefaultConfigFileName is the default name for the chezmoi config file.
// Prefer TOML for its readability and comment support.
const DefaultConfigFileName = "chezmoi.toml"

// SupportedConfigFileNames lists all config file names chezmoi recognizes.
// Order matters: the first match wins during config file discovery.
var SupportedConfigFileNames = []string{
	"chezmoi.toml",
	"chezmoi.yaml",
	"chezmoi.json",
}

// SourcePrefix is the prefix used for source-specific files and directories.
const SourcePrefix = "dot_"

// PrivatePrefix is the prefix for private files (mode 0600).
const PrivatePrefix = "private_"

// ExecutablePrefix is the prefix for executable files (mode 0755).
const ExecutablePrefix = "executable_"

// EncryptedSuffix is the suffix for encrypted files.
const EncryptedSuffix = ".age"

// TemplateSuffix is the suffix for template files.
const TemplateSuffix = ".tmpl"

// UserHomeDir returns the current user's home directory, respecting
// the HOME environment variable on Unix systems.
func UserHomeDir() (string, error) {
	if home := os.Getenv("HOME"); home != "" && runtime.GOOS != "windows" {
		return home, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("could not determine home directory: " + err.Error())
	}
	return home, nil
}

// DefaultSourceDir returns the default source directory path for the current user.
func DefaultSourceDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", errors.New("could not determine config directory: " + err.Error())
	}
	return filepath.Join(configDir, DefaultSourceDirName), nil
}

// DefaultConfigFile returns the default config file path for the current user.
func DefaultConfigFile() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", errors.New("could not determine config directory: " + err.Error())
	}
	return filepath.Join(configDir, DefaultSourceDirName, DefaultConfigFileName), nil
}

// FileMode represents the permission mode of a file managed by chezmoi.
type FileMode uint32

const (
	// FileModeRegular is the default mode for regular files.
	FileModeRegular FileMode = 0o644
	// FileModePrivate is the mode for private files (e.g. SSH keys, credentials).
	FileModePrivate FileMode = 0o600
	// FileModeExecutable is the mode for executable files.
	FileModeExecutable FileMode = 0o755
	// FileModeDir is the default mode for directories.
	FileModeDir FileMode = 0o755
	// FileModeDirPrivate is the mode for private directories.
	FileModeDirPrivate FileMode = 0o700
)
