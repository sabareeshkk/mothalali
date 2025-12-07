package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func UpdateRef(ref string, oid string) error {
	filePath := GitDir + "/" + ref
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	err := os.WriteFile(filePath, []byte(oid), 0644) // TODO: add meaningful name to variable 0644
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

func GetRef(ref string) (string, error) {
	filePath := GitDir + "/" + ref
	// Check if file exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	// Strip whitespace (newline, spaces, etc.)
	oid := strings.TrimSpace(string(content))
	return oid, nil
}
