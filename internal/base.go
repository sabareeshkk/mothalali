package internal

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func createHashObject(content []byte, objType string) (oid []byte, err error) {
	var header []byte
	header = fmt.Appendf(header, "%s %d\x00", objType, len(content))
	fileContent := append(header, content...)
	hash := sha1.New()
	hash.Write(fileContent)
	oid = hash.Sum(nil)
	dir := filepath.Join(ObjectsDir, fmt.Sprintf("%x", oid[:1]))
	filePath := filepath.Join(dir, fmt.Sprintf("%x", oid[1:]))
	// Ensure the directory exists
	if err := os.MkdirAll(dir, 0755); err != nil { // TODO: add meaningful name to variable 0755
		fmt.Println("Error creating directory:", err)
		return nil, err
	}
	err = os.WriteFile(filePath, fileContent, 0644) // TODO: add meaningful name to variable 0644
	if err != nil {
		fmt.Println("Error writing file:", err)
		return nil, err
	}
	fmt.Printf("file_path: %s\n", filePath)
	return oid, nil
}

func HashObject(path string, obj_type string) (string, error) {
	var content []byte
	switch obj_type {
	case "blob":
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("File not found:", path)
			return "", err
		}
		var err error
		content, err = os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return "", err
		}
	case "tree":
		content = []byte(path) // Converts string to byte slice
	default:
		content = []byte(path) // Converts string to byte slice
	}
	oid, err := createHashObject(content, obj_type)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", oid), nil
}

func parseGitObject(object []byte) (objType string, size int, content []byte, err error) {
	spaceIdx := bytes.IndexByte(object, ' ')
	nullIdx := bytes.IndexByte(object, 0)
	if spaceIdx == -1 || nullIdx == -1 || nullIdx < spaceIdx {
		err = fmt.Errorf("invalid git object format")
		return
	}

	objType = string(object[:spaceIdx])
	sizeStr := string(object[spaceIdx+1 : nullIdx])
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		return
	}
	content = object[nullIdx+1:]
	return
}

func ReadObject(sha1_hash string, expected string) ([]byte, error) {
	dir := filepath.Join(ObjectsDir, sha1_hash[:2])
	filepath := filepath.Join(dir, sha1_hash[2:])
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	objType, size, content, err := parseGitObject(content)
	fmt.Println("objType:", objType, "size:", size)
	if err != nil {
		fmt.Println("Error parsing git object:", err)
		return nil, err
	}
	if expected != "" && objType != expected {
		return nil, fmt.Errorf("Unexpected object type: %s", objType)
	}
	return content, nil
}

// Implement logic to check if the path should be ignored
func ignoredPath(path string) bool {
	paths := []string{".git", ".mothalali", "build"}
	if slices.Contains(paths, path) {
		return true
	}
	return false
}

// Each entry is [name, oid, objType]
func SortAndJoinEntries(entries [][]string) string {
	// Sort by name (first element)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i][0] < entries[j][0]
	})

	var builder strings.Builder
	for _, entry := range entries {
		name, oid, objType := entry[0], entry[1], entry[2]
		builder.WriteString(fmt.Sprintf("%s %s %s\n", objType, oid, name))
	}
	return builder.String()
}

func WriteTree(directoryPath string) (string, error) {
	var dirEntries [][]string
	// Create a new tree object
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		fullPath := fmt.Sprintf("%s/%s", directoryPath, entry.Name())
		if ignoredPath(entry.Name()) { // check if the path is ignored
			continue
		}
		info, err := entry.Info()
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}
		if info.IsDir() && (info.Mode()&os.ModeSymlink == 0) { // This is a real directory, not a symlink
			fmt.Printf("\033[32m Directory -- %s \n", entry.Name())
			obj_type := "tree"
			oid, err := WriteTree(fullPath)
			if err != nil {
				fmt.Println("Error writing tree:", err)
			} else {
				fmt.Println("successfully created:", oid)
			}
			dirEntries = append(dirEntries, []string{entry.Name(), oid, obj_type})
		} else if info.Mode().IsRegular() { // This is a regular file, not a symlink
			fmt.Printf("\033[31m Writing path %s\n", fullPath)
			obj_type := "blob" // TODO: move to internal/constants.go
			oid, err := HashObject(fullPath, obj_type)
			if err != nil {
				fmt.Println("Error hashing object:", err)
			} else {
				fmt.Println("successfully created:", oid)
			}
			dirEntries = append(dirEntries, []string{entry.Name(), oid, obj_type})
		}
	}
	tree := SortAndJoinEntries(dirEntries)
	fmt.Println("Tree:", tree)
	oid, err := HashObject(tree, "tree")
	if err != nil {
		fmt.Println("Error hashing tree:", err)
	}
	return oid, nil
}

func parseTree(oid string) <-chan struct {
	Type, Oid, Name string
} {
	ch := make(chan struct{ Type, Oid, Name string })
	go func() {
		defer close(ch)
		if oid == "" {
			return
		}

		data, err := ReadObject(oid, "tree")
		if err != nil {
			return
		}

		for line := range strings.SplitSeq(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Fields(line)
			if len(parts) < 3 {
				continue
			}
			ch <- struct{ Type, Oid, Name string }{parts[0], parts[1], strings.Join(parts[2:], " ")}
		}
	}()

	return ch
}

func getTree(oid string, basePath string) (index map[string]string) {
	index = make(map[string]string)
	getTreeInto(oid, basePath, index)
	return index
}

func getTreeInto(oid string, basePath string, index map[string]string) (treePath map[string]string) {
	treePath = make(map[string]string)
	for entry := range parseTree(oid) {
		fullPath := filepath.Join(basePath, entry.Name)
		fmt.Println("fullPath:", fullPath, basePath)
		if entry.Type == "tree" {
			getTreeInto(entry.Oid, fullPath, index)
		} else if entry.Type == "blob" {
			index[fullPath] = entry.Oid
		}
	}
	return treePath
}

func ReadTree(oid string) error {
	treePath := getTree(oid, "./")
	for path, oid := range treePath {
		fmt.Println("Exact path:", path, oid)
		content, err := ReadObject(oid, "blob")
		if err != nil {
			return err
		}
		fmt.Println("Content:", string(content))
	}
	return nil
}
