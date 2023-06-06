package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func runFormatter(inputCode string) ([]byte, error) {
	cmd := exec.Command("rustfmt")
	cmd.Stdin = strings.NewReader(inputCode)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
