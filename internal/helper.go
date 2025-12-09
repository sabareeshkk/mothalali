package internal

import (
	"bytes"
	"os"
	"os/exec"
)

func ShowDot(dot string) error {

    cmd := exec.Command("dot", "-Tpng")
    cmd.Stdin = bytes.NewBufferString(dot)
    var buf bytes.Buffer
    cmd.Stdout = &buf

    err := cmd.Run()
    if err != nil {
        panic(err)
    }
	return os.WriteFile("graph.png", buf.Bytes(), 0644)
}
