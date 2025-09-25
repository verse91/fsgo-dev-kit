package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

// CreateFile creates a file with the given content at the specified path.
// It creates the directory structure if it doesn't exist.
func CreateFile(path, content string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// CreateDirectory creates a directory with the given permissions
func CreateDirectory(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// RunCommand executes a shell command with stdout and stderr connected
func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunCommandInDir executes a shell command in a specific directory
func RunCommandInDir(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GetCurrentDir returns the current working directory
func GetCurrentDir() (string, error) {
	return os.Getwd()
}

// ChangeDir changes the current working directory
func ChangeDir(dir string) error {
	return os.Chdir(dir)
}

// GetProjectName returns the base name of the current directory as project name
func GetProjectName() (string, error) {
	currentDir, err := GetCurrentDir()
	if err != nil {
		return "", err
	}
	return filepath.Base(currentDir), nil
}