package internal

import (
	"fmt"
	"os"
	"strings"
)

func SetHead(oid string) error {
	err := os.WriteFile(HeadFile, []byte(oid), 0644) // TODO: add meaningful name to variable 0644
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

func GetHead() (string, error) {
	// Check if file exists
	if _, err := os.Stat(HeadFile); err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	content, err := os.ReadFile(HeadFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	// Strip whitespace (newline, spaces, etc.)
	oid := strings.TrimSpace(string(content))
	return oid, nil
}
