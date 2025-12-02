package internal

import(
	"fmt"
	"os"
)

func SetHead(oid string) error {
	err := os.WriteFile(HeadDir, []byte(oid), 0644) // TODO: add meaningful name to variable 0644
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}
