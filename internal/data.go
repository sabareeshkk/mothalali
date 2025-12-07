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

type Ref struct {
	Name string
	OID  string
}

func IterRefs() <-chan Ref {
	ch := make(chan Ref)

	go func() {
		defer close(ch)

		// HEAD
		if oid, err := GetRef("HEAD"); err == nil {
			ch <- Ref{Name: "HEAD", OID: oid}
		}

		// refs/*
		root := filepath.Join(GitDir, "refs")
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			rel, err := filepath.Rel(GitDir, path)
			if err != nil {
				return err
			}
			oid, err := GetRef(rel)
			if err != nil {
				return err
			}
			ch <- Ref{Name: rel, OID: oid}
			return nil
		})

	}()

	return ch
}
